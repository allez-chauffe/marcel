import { actions } from './actions'

const intialState = { plugins: '', props: '', clients: '' }

const filters = (state = intialState, action) => {
  switch (action.type) {
    case actions.CHANGE_FILTER:
      return { ...state, [action.payload.collection]: action.payload.filter }
    default:
      return state
  }
}

export default filters
