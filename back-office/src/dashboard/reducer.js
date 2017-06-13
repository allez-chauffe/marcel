//@flow
import type { Reducer } from 'redux'
import { actions } from './actions'
import type { DashboardAction, DashboardState } from './type'

const intialState = {
  selectedPlugin: null,
}

const dashboard: Reducer<DashboardState, DashboardAction> = (
  state = intialState,
  action,
) => {
  switch (action.type) {
    case actions.SELECT_PLUGIN:
      return { ...state, selectedPlugin: action.payload.elementName }
    default:
      return state
  }
}

export default dashboard
