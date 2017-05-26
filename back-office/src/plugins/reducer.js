// @flow
import { range } from 'lodash'
import type { Reducer } from 'redux'
import type { State } from './plugins.types'
import type { Action } from '../store/types'

const intialState = range(20).map(i => ({
  name: `Plugin ${i}`,
  elementName: `plugin-${i}`,
  icon: 'picture_in_picture_alt',
}))

const reducer: Reducer<State, Action> = (state = intialState, action) => {
  return state
}

export default reducer
