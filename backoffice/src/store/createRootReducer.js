import { combineReducers } from 'redux'
import { reducer as toastr } from 'react-redux-toastr'
import { connectRouter } from 'connected-react-router'

import { reducer as plugins } from '../plugins'
import filters from './filters'
import { reducer as dashboard } from '../dashboard'
import { reducer as auth } from '../auth'
import { reducer as clients } from '../clients'
import { reducer as loaders, actions as loadersActions } from './loaders'
import { reducer as users } from '../user'

const uris = (state = { API: '/api/' }, action) => {
  switch (action.type) {
    case loadersActions.LOAD_URIS_SUCCESSED: {
      return action.payload.uris
    }
    default:
      return state
  }
}

const createRootReducer = history => combineReducers({
  router: connectRouter(history),
  plugins,
  filters,
  dashboard,
  clients,
  toastr,
  auth,
  loaders,
  uris,
  users,
})

export default createRootReducer
