// @flow
import { combineReducers } from 'redux'
import type { Reducer } from 'redux'

import { reducer as toastr } from 'react-redux-toastr'

import { reducer as plugins } from '../plugins'
import filters from './filters'
import { reducer as dashboard } from '../dashboard'
import { reducer as auth } from '../auth'
import { reducer as clients } from '../clients'
import { reducer as loaders, actions as loadersActions } from './loaders'
import type { State, Action, Config } from './types'

const config: Reducer<Config, Action> = (
  state = { backendURI: 'http://localhost:8090/api/v1/', frontendURI: 'http://localhost:5000/' },
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

const rootReducer: Reducer<State, Action> = combineReducers({
  plugins,
  filters,
  dashboard,
  clients,
  toastr,
  auth,
  loaders,
  config,
})

export default rootReducer
