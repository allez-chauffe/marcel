// @flow
import { range } from 'lodash'
import type { Reducer } from 'redux'
import type { State } from './plugins.type'
import type { Action } from '../store/types'

const intialState = range(20).map(i => ({
  name: `Plugin ${i}`,
  elementName: `plugin-${i}`,
  icon: 'picture_in_picture_alt',
  props: [
    { name: 'prop1', type: 'string', value: 'hello world !' },
    { name: 'prop2', type: 'number', value: 42 },
    { name: 'prop3', type: 'boolean', value: true },
    { name: 'prop4', type: 'json', value: { collection: ['first', 'second'] } },
  ],
}))

const reducer: Reducer<State, Action> = (state = intialState, action) => state

export default reducer
