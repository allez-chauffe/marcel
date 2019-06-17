import { mapValues, keyBy, keys, chain as _chain, pickBy, omit, map, flatten, reduce } from 'lodash'
import { set, unset, chain, update } from 'immutadot'
import { actions } from './actions'
import { actions as loadActions } from '../store/loaders'
import { values } from 'lodash/fp'
import uuid from 'uuid/v4'
import { getPluginInstances } from '../common/utils'

const intialState = {
  selectedPlugin: null,
  selectedDashboard: null,
  deletingDashboard: null,
  loading: false,
  dashboards: {},
  pluginInstances: {},
}

const updatePlugins = (layout) => (plugins) => {
  const updatedInstances = mapValues(layout, (layoutItem, instanceId) => {
    const plugin = plugins[instanceId]
    if (!plugin) throw new Error('Plugin instance not found in layout')

    const { x, y, w: cols, h: rows } = layoutItem
    return { ...plugin, x, y, cols, rows }
  })

  return { ...plugins, ...updatedInstances }
}

const dashboard = (state = intialState, action) => {
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
    case actions.DASHBOARD_DELETED: {
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
        .set(`dashboards.${dashboard.id}`, dashboard)
        .set('selectedDashboard', dashboard.id)
        .value()
    }
    case actions.ADD_SUB_PLUGIN: {
      const { selectedPlugin } = state
      if (!selectedPlugin) return state

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
        .push(`pluginInstances.${selectedPlugin}.props.${propName}.value`, instanceId)
        .value()
    }
    case actions.REORDER_SUB_PLUGINS: {
      const { instanceIds, parent: { plugin, prop } } = action.payload
      return set(state, `pluginInstances.${plugin}.props.${prop}.value`, instanceIds)
    }
    case actions.ADD_PLUGIN: {
      const instanceId = uuid()
      const { selectedDashboard } = state
      return selectedDashboard
        ? chain(state)
            .set('selectedPlugin', instanceId)
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
      const { plugin } = action.payload
      const { selectedDashboard } = state
      if (!selectedDashboard) return state

      const { instanceId, parent } = plugin

      const removeChilds = (pluginInstances, instanceId) => {
        const plugin = pluginInstances[instanceId]
        if (!plugin) return pluginInstances

        const pluginListProps = pickBy(plugin.props, { type: 'pluginList' })
        const pluginsToRemove = flatten(map(values(pluginListProps), 'value'))

        if (!pluginsToRemove.lenght) return pluginInstances

        const cleanedPluginInstances = omit(pluginInstances, pluginsToRemove)
        return reduce(pluginsToRemove, removeChilds, cleanedPluginInstances)
      }

      return chain(state)
        .update(`pluginInstances`, pluginInstances => removeChilds(pluginInstances, instanceId))
        .unset(`pluginInstances.${instanceId}`)
        .pull(
          parent
            ? `pluginInstances.${parent.plugin}.props.${parent.prop}.value`
            : `dashboards.${selectedDashboard}.plugins`,
          instanceId,
        )
        .set(`selectedPlugin`, parent && parent.plugin)
        .value()
    }
    case actions.CHANGE_PROP: {
      const { instanceId, prop, value } = action.payload
      return set(state, `pluginInstances.${instanceId}.props.${prop.name}.value`, value)
    }
    case actions.SAVE_LAYOUT: {
      const { layout } = action.payload
      const { selectedDashboard } = state
      return selectedDashboard ? update(state, `pluginInstances`, updatePlugins(layout)) : state
    }
    case actions.UPDATE_CONFIG: {
      const { selectedDashboard } = state
      const { property, value } = action.payload
      const parsedValue = parseFloat(value)
      return selectedDashboard
        ? set(
            state,
            `dashboards.${selectedDashboard}.${property}`,
            isNaN(parsedValue) ? value : parsedValue,
          )
        : state
    }
    case loadActions.LOAD_DASHBOARDS_SUCCESSED: {
      const { dashboards, user } = action.payload
      const plugins = _chain(dashboards)
        .map('plugins')
        .map(values)
        .flatten()
        .value()
      const pluginInstances = getPluginInstances(plugins)
      const normalizedDashboards = dashboards.map(dashboard => ({
        ...dashboard,
        isWritable: user.role === "admin" || user.id === dashboard.owner,
        plugins: keys(dashboard.plugins),
      }))

      return {
        ...state,
        dashboards: keyBy(normalizedDashboards, 'id'),
        pluginInstances: keyBy(pluginInstances, 'instanceId'),
      }
    }
    case actions.SELECT_PLUGIN_PARENT: {
      const { selectedPlugin, pluginInstances } = state
      if (!selectedPlugin) return state

      const plugin = pluginInstances[selectedPlugin]
      if (!plugin) return state

      const { parent } = plugin
      return parent ? { ...state, selectedPlugin: parent.plugin } : state
    }
    case actions.ACTIVATE_DASHBOARD: {
      const { dashboardId } = action.payload
      return set(state, `dashboards.${dashboardId}.isactive`, true)
    }
    case actions.DEACTIVATE_DASHBOARD: {
      const { dashboardId } = action.payload
      return set(state, `dashboards.${dashboardId}.isactive`, false)
    }
    default:
      return state
  }
}

export default dashboard
