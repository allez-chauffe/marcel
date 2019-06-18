import { combineReducers } from 'redux'

import { reducer as toastr } from 'react-redux-toastr'

import { reducer as plugins } from '../plugins'
import filters from './filters'
import { reducer as dashboard } from '../dashboard'
import { reducer as auth } from '../auth'
import { reducer as clients } from '../clients'
import { reducer as loaders, actions as loadersActions } from './loaders'
import { reducer as users } from '../user'

const config = (
  state = {
    backendURI: 'http://marcel.com:8081/api/v1/',
    authURI: 'http://marcel.com:8081/auth/',
    frontendURI: 'http://marcel.com:8081/front/',
  },
  action,
) => {
  switch (action.type) {
    case loadersActions.LOAD_CONFIG_SUCCESSED: {
      return action.payload.config
    }
    default:
      return state
  }
}

const rootReducer = combineReducers({
  plugins,
  filters,
  dashboard,
  clients,
  toastr,
  auth,
  loaders,
  config,
  users,
})

export default rootReducer
