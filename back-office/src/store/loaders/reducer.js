//@flow
import type { Reducer } from 'redux'
import type { LoadersState, LoadersAction } from './type'
import { actions } from './actions'

const initialState = {
  config: true,
  plugins: false,
  dashboards: false,
  clients: false,
  initial: false,
}

const loaders: Reducer<LoadersState, LoadersAction> = (state = initialState, action) => {
  switch (action.type) {
    case actions.LOAD_INITIAL_STARTED:
      return { ...state, initial: true }
    case actions.LOAD_INITIAL_FINISHED:
      return { ...state, initial: false }
    case actions.LOAD_CONFIG_STARTED:
      return { ...state, config: true }
    case actions.LOAD_CONFIG_SUCCESSED:
    case actions.LOAD_CONFIG_FAILED:
      return { ...state, config: false }
    case actions.LOAD_DASHBOARDS_STARTED:
      return { ...state, dashboards: true }
    case actions.LOAD_DASHBOARDS_SUCCESSED:
    case actions.LOAD_DASHBOARDS_FAILED:
      return { ...state, dashboards: false }
    case actions.LOAD_PLUGINS_STARTED:
      return { ...state, plugins: true }
    case actions.LOAD_PLUGINS_SUCCESSED:
    case actions.LOAD_PLUGINS_FAILED:
      return { ...state, plugins: false }
    default:
      return state
  }
}

export default loaders
