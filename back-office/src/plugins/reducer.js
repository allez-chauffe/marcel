import { actions as loadActions } from '../store/loaders'

const intialState = []

const reducer = (state = intialState, action) => {
  switch (action.type) {
    case loadActions.LOAD_PLUGINS_SUCCESSED:
      return action.payload.plugins
    default:
      return state
  }
}

export default reducer
