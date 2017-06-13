// @flow
import { combineReducers } from 'redux'
import type { Reducer } from 'redux'

import { reducer as plugins } from '../plugins'
import filters from './filters'
import type { State, Action } from './types'

const rootReducer: Reducer<State, Action> = combineReducers({
  plugins,
  filters,
})

export default rootReducer
