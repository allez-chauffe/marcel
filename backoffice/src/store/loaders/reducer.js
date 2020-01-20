import { actions } from './actions'

const initialState = {
  uris: true,
  plugins: false,
  dashboards: false,
  clients: false,
  initial: false,
  users: false,
}

const loaders = (state = initialState, action) => {
  switch (action.type) {
    case actions.LOAD_INITIAL_STARTED:
      return { ...state, initial: true }
    case actions.LOAD_INITIAL_FINISHED:
      return { ...state, initial: false }
    case actions.LOAD_URIS_STARTED:
      return { ...state, uris: true }
    case actions.LOAD_URIS_SUCCESSED:
    case actions.LOAD_URIS_FAILED:
      return { ...state, uris: false }
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
    case actions.LOAD_USERS_STARTED:
      return { ...state, users: true }
    case actions.LOAD_USERS_SUCCESSED:
    case actions.LOAD_USERS_FAILED:
      return { ...state, users: false }
    default:
      return state
  }
}

export default loaders
