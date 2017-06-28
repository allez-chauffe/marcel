//@flow
import { combineReducers } from 'redux'
import type { Reducer } from 'redux'
import type { AuthState, AuthAction } from './type'
import { actions } from './actions'

const token: Reducer<$PropertyType<AuthState, 'token'>, AuthAction> = (
  state = null,
  action,
) => {
  switch (action.type) {
    case actions.LOGIN_SUCCESS: {
      return action.payload.token
    }
    case actions.LOGIN_FAIL:
    case actions.LOGOUT: {
      return null
    }
    default: {
      return state
    }
  }
}

const initialForm = { login: '', password: '' }
const form: Reducer<$PropertyType<AuthState, 'form'>, AuthAction> = (
  state = initialForm,
  action,
) => {
  switch (action.type) {
    case actions.CHANGE_LOGIN: {
      const { login } = action.payload
      return { ...state, login }
    }
    case actions.CHANGE_PASSWORD: {
      const { password } = action.payload
      return { ...state, password }
    }
    case actions.RESET_FORM: {
      return initialForm
    }
    default:
      return state
  }
}

const auth: Reducer<AuthState, AuthAction> = combineReducers({
  token,
  form,
})

export default auth
