// @flow
import type { Reducer } from 'redux'
import type { State } from './type'
import type { Action } from '../store/types'
import { actions as loadActions } from '../store/loaders'

const intialState = []

const reducer: Reducer<State, Action> = (state = intialState, action) => {
  switch (action.type) {
    case loadActions.LOAD_PLUGINS_SUCCESSED:
      return action.payload.plugins
    default:
      return state
  }
}

export default reducer
