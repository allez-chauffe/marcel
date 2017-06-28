// @flow
import { combineReducers } from 'redux'
import type { Reducer } from 'redux'

import { reducer as toastr } from 'react-redux-toastr'

import { reducer as plugins } from '../plugins'
import filters from './filters'
import { reducer as dashboard } from '../dashboard'
import { reducer as auth } from '../auth'
import type { State, Action } from './types'

const rootReducer: Reducer<State, Action> = combineReducers({
  plugins,
  filters,
  dashboard,
  toastr,
  auth,
})

export default rootReducer
