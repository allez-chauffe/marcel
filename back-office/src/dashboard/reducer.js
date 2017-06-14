//@flow
import type { Reducer } from 'redux'
import { actions } from './actions'
import { push, map } from 'immutadot'
import { find } from 'lodash'
import type {
  DashboardAction,
  DashboardState,
  Dashboard,
  LayoutMap,
  PluginInstance,
} from './type'
import type { Plugin } from '../plugins'
import uuid from 'uuid/v4'

const intialState = {
  selectedPlugin: null,
  dashboard: {
    name: 'Dashboard',
    description: 'Some description',
    plugins: [
      {
        name: `Plugin 1`,
        elementName: `plugin-1`,
        instanceId: 'plugin-1#0',
        icon: 'picture_in_picture_alt',
        x: 0,
        y: 0,
        columns: 2,
        rows: 3,
        props: [],
      },
    ],
  },
}

const updatePlugin = (layout: LayoutMap) => (plugin: PluginInstance) => {
  if (!layout[plugin.instanceId])
    throw new Error('Plugin instance not foun in layout')

  const { x, y, w: columns, h: rows } = layout[plugin.instanceId]
  return { ...plugin, x, y, columns, rows }
}

const dashboard: Reducer<DashboardState, DashboardAction> = (
  state = intialState,
  action,
) => {
  switch (action.type) {
    case actions.SELECT_PLUGIN:
      return { ...state, selectedPlugin: action.payload.instanceId }
    case actions.ADD_PLUGIN:
      return push(state, 'dashboard.plugins', {
        ...action.payload.plugin,
        x: 0,
        y: 0,
        columns: 1,
        rows: 1,
        instanceId: uuid(),
      })
    case actions.SAVE_LAYOUT:
      const { layout } = action.payload
      return map(state, 'dashboard.plugins', updatePlugin(layout))
    default:
      return state
  }
}

export default dashboard
