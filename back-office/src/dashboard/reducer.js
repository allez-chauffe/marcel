//@flow
import type { Reducer } from 'redux'
import { actions } from './actions'
import { actions as loadActions } from '../store/loaders'
import { mapValues, keyBy, keys, chain as _chain } from 'lodash'
import { values } from 'lodash/fp'
import { set, unset, chain, update } from 'immutadot'
import uuid from 'uuid/v4'
import { getPluginInstances } from '../common/utils'
import type {
  DashboardAction,
  DashboardState,
  LayoutMap,
  PluginInstanceMap,
} from './type'

const intialState = {
  selectedPlugin: null,
  selectedDashboard: null,
  deletingDashboard: null,
  displayGrid: true,
  loading: false,
  dashboards: {},
  pluginInstances: {},
}

const updatePlugins = (layout: LayoutMap) => (plugins: PluginInstanceMap) => {
  return mapValues(layout, (layoutItem, instanceId) => {
    const plugin = plugins[instanceId]
    if (!plugin) throw new Error('Plugin instance not found in layout')

    const { x, y, w: cols, h: rows } = layoutItem
    return { ...plugin, x, y, cols, rows }
  })
}

const dashboard: Reducer<DashboardState, DashboardAction> = (
  state = intialState,
  action,
) => {
  switch (action.type) {
    case actions.SELECT_PLUGIN: {
      return { ...state, selectedPlugin: action.payload.instanceId }
    }
    case actions.SELECT_DASHBOARD: {
      return { ...state, selectedDashboard: action.payload.dashboardId }
    }
    case actions.UNSELECT_DASHBOARD: {
      return { ...state, selectedDashboard: null, selectedPlugin: null }
    }
    case actions.REQUIRE_DASHBOARD_DELETION: {
      return { ...state, deletingDashboard: action.payload.dashboardId }
    }
    case actions.CONFIRM_DASHBOARD_DELETION: {
      const { deletingDashboard } = state
      return deletingDashboard
        ? chain(state)
            .set('deletingDashboard', null)
            .unset(`dashboards.${deletingDashboard}`)
            .value()
        : { ...state, deletingDashboard: null }
    }
    case actions.CANCEL_DASHBOARD_DELETION: {
      return { ...state, deletingDashboard: null }
    }
    case actions.DELETE_DASHBOARD: {
      return unset(state, `dashboards.${action.payload.dashboardId}`)
    }
    case actions.ADD_DASHBOARD: {
      const { dashboard } = action.payload
      return chain(state)
        .set(`dashboards.${dashboard.id}`, {
          ...dashboard,
          name: `Dashboard ${dashboard.id}`,
        })
        .set('selectedDashboard', dashboard.id)
        .value()
    }
    case actions.ADD_SUB_PLUGIN: {
      const { selectedPlugin } = state
      const { propName, plugin } = action.payload
      const instanceId = uuid()
      return chain(state)
        .set(`pluginInstances.${instanceId}`, {
          ...plugin,
          x: 0,
          y: 0,
          cols: 1,
          rows: 1,
          instanceId,
          parent: { plugin: selectedPlugin, prop: propName },
        })
        .push(
          `pluginInstances.${selectedPlugin}.props.${propName}.value`,
          instanceId,
        )
        .value()
    }
    case actions.ADD_PLUGIN: {
      const instanceId = uuid()
      const { selectedDashboard } = state
      return selectedDashboard
        ? chain(state)
            .set(`pluginInstances.${instanceId}`, {
              ...action.payload.plugin,
              x: action.payload.x,
              y: action.payload.y,
              cols: 1,
              rows: 1,
              instanceId,
            })
            .push(`dashboards.${selectedDashboard}.plugins`, instanceId)
            .value()
        : state
    }
    case actions.DELETE_PLUGIN: {
      const { selectedPlugin, selectedDashboard } = state
      if (!selectedPlugin || !selectedDashboard) return state
      return chain(state)
        .unset(`plugins.${selectedPlugin}`)
        .pull(`dashboards.${selectedDashboard}.plugins`, selectedPlugin)
    }
    case actions.CHANGE_PROP: {
      const { instanceId, prop, value } = action.payload
      return set(state, `plugins.${instanceId}.props.${prop.name}.value`, value)
    }
    case actions.SAVE_LAYOUT: {
      const { layout } = action.payload
      const { selectedDashboard } = state
      return selectedDashboard
        ? update(state, `pluginInstances`, updatePlugins(layout))
        : state
    }
    case actions.UPDATE_CONFIG: {
      const { selectedDashboard } = state
      const { property, value } = action.payload
      const parsedValue = !isNaN(value) ? parseFloat(value) : value
      return selectedDashboard
        ? set(state, `dashboards.${selectedDashboard}.${property}`, parsedValue)
        : state
    }
    case actions.TOGGLE_DISPLAY_GRID: {
      return { ...state, displayGrid: !state.displayGrid }
    }
    case loadActions.LOAD_DASHBOARDS_SUCCESSED: {
      const { dashboards } = action.payload
      const plugins = _chain(dashboards)
        .map('plugins')
        .map(values)
        .flatten()
        .value()
      const pluginInstances = getPluginInstances(plugins)
      const normalizedDashboards = dashboards.map(dashboard => ({
        ...dashboard,
        plugins: keys(dashboard.plugins),
      }))

      return {
        ...state,
        dashboards: keyBy(normalizedDashboards, 'id'),
        pluginInstances: keyBy(pluginInstances, 'instanceId'),
      }
    }
    default:
      return state
  }
}

export default dashboard
