// @flow
import { combineReducers } from 'redux'
import type { Reducer } from 'redux'

import { reducer as plugins } from '../plugins'
import type { State, Action } from './types'

const rootReducer: Reducer<State, Action> = combineReducers({
  plugins,
})

export default rootReducer
