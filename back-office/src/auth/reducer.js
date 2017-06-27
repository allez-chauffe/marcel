//@flow
import type { Reducer } from 'redux'
import type { AuthState, AuthAction } from './type'
import { actions } from './actions'

export const initialState = {
  token: null,
}

const auth: Reducer<AuthState, AuthAction> = (state = initialState, action) => {
  switch (action.type) {
    case actions.LOGIN_SUCCESS: {
      return { ...state, token: action.payload.token }
    }
    case actions.LOGIN_FAIL:
    case actions.LOGOUT: {
      return { ...state, token: null }
    }
    default: {
      return state
    }
  }
}

export default auth
