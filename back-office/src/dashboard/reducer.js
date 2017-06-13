//@flow
import type { Reducer } from 'redux'
import { actions } from './actions'
import { push } from 'immutadot'
import type { DashboardAction, DashboardState } from './type'

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
        instanceId: 'lol',
      })

    default:
      return state
  }
}

export default dashboard
