// @flow
import { combineReducers } from 'redux'
import type { Reducer } from 'redux'

import { reducer as plugins } from '../plugins'
import filters from './filters'
import { reducer as dashboard } from '../dashboard'
import type { State, Action } from './types'

const rootReducer: Reducer<State, Action> = combineReducers({
  plugins,
  filters,
  dashboard,
})

export default rootReducer
