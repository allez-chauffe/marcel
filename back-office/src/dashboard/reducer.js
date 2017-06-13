//@flow
import type { Reducer } from 'redux'
import { actions } from './actions'
import { push } from 'immutadot'
import { find } from 'lodash'
import type { DashboardAction, DashboardState, Dashboard } from './type'
import type { Plugin } from '../plugins'

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

const generateInstanceId = (dashboard: Dashboard, plugin: Plugin) => {
  const { elementName } = plugin
  let id = 0
  while (find(dashboard.plugins, { instanceId: `${elementName}#${id}` }))
    id++
  return `${elementName}#${id}`
}

const dashboard: Reducer<DashboardState, DashboardAction> = (
  state = intialState,
  action,
) => {
  console.log(state)
  switch (action.type) {
    case actions.SELECT_PLUGIN:
      return { ...state, selectedPlugin: action.payload.instanceId }
    case actions.ADD_PLUGIN:
      console.log(state)
      return push(state, 'dashboard.plugins', {
        ...action.payload.plugin,
        x: 0,
        y: 0,
        columns: 1,
        rows: 1,
        instanceId: generateInstanceId(state.dashboard, action.payload.plugin),
      })

    default:
      return state
  }
}

export default dashboard
