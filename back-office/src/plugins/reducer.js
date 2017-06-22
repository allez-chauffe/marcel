// @flow
import { range } from 'lodash'
import type { Reducer } from 'redux'
import type { State } from './type'
import type { Action } from '../store/types'

import mockedData from '../mocked-data/plugins'
const intialState = mockedData

const reducer: Reducer<State, Action> = (state = intialState, action) => state

export default reducer
