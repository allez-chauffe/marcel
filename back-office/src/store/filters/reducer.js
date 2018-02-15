// @flow
import type { Reducer } from 'redux'
import type { FiltersState, FiltersAction } from './types'
import { actions } from './actions'

const intialState = { plugins: '', props: '', clients: '' }

const filters: Reducer<FiltersState, FiltersAction> = (
  state = intialState,
  action,
) => {
  switch (action.type) {
    case actions.CHANGE_FILTER:
      return { ...state, [action.payload.collection]: action.payload.filter }
    default:
      return state
  }
}

export default filters
