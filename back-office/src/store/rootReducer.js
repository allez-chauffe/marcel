// @flow
import { combineReducers } from 'redux'
import type { Reducer } from 'redux'

import { reducer as toastr } from 'react-redux-toastr'

import { reducer as plugins } from '../plugins'
import filters from './filters'
import { reducer as dashboard } from '../dashboard'
import { reducer as auth } from '../auth'
import { reducer as loaders, actions as loadersActions } from './loaders'
import type { State, Action, Config } from './types'

const config: Reducer<Config, Action> = (
  state = { backendURI: '' },
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
  toastr,
  auth,
  loaders,
  config,
})

export default rootReducer
