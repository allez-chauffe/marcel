import { set, chain, unset } from 'immutadot'
import { actions as loadersActions } from '../store/loaders'
import { actions } from './actions'

export const initialState = {
  clients: {},
  loading: {},
  associating: null,
}

const clients = (state = initialState, action) => {
  switch (action.type) {
    case loadersActions.LOAD_CLIENTS_SUCCESSED: {
      return { ...state, clients: action.payload.clients }
    }
    case actions.CLIENT_ASSOCIATION_STARTED: {
      return set(state, `loading.${action.payload.client.id}`, true)
    }
    case actions.CLIENT_ASSOCIATION_SUCCESS: {
      const { client } = action.payload
      return chain(state)
        .unset(`loading.${client.id}`)
        .set(`clients.${client.id}`, client)
        .value()
    }
    case actions.CLIENT_ASSOCIATION_FAILED: {
      const { client } = action.payload
      return unset(state, `loading.${client.id}`)
    }
    case actions.REQUIRE_CLIENT_ASSOCIATION: {
      const { client } = action.payload
      return { ...state, associating: client }
    }
    case actions.CONFIRM_CLIENT_ASSOCIATION:
    case actions.CANCEL_CLIENT_ASSOCIATION: {
      return { ...state, associating: null }
    }
    default:
      return state
  }
}

export default clients
