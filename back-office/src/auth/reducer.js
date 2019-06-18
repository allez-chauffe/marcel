import { combineReducers } from 'redux'
import { actions } from './actions'
import { set } from 'immutadot'

const user = (state = null, action) => {
  switch (action.type) {
    case actions.LOGIN_SUCCESS: {
      return action.payload.user
    }
    case actions.LOGOUT_SUCCESS:
    case actions.LOGIN_FAIL:
    case actions.DISCONNECTED: {
      return null
    }
    default: {
      return state
    }
  }
}

const isLoading = (state = false, action) => {
  if (!action.type.startsWith('AUTH/')) return state

  if (action.type.endsWith('REQUEST')) return true
  if (action.type.endsWith('SUCCESS') || action.type.endsWith('FAIL')) return false

  return state
}

const initialForm = { login: '', password: '' }
const form = (state = initialForm, action) => {
  switch (action.type) {
    case actions.CHANGE_LOGIN: {
      return { ...state, login: action.payload.login }
    }
    case actions.CHANGE_PASSWORD: {
      return { ...state, password: action.payload.password }
    }
    case actions.RESET_FORM: {
      return initialForm
    }
    case actions.UPDATE_CONNECTED_USER_SUCCESS: {
      const { user } = action.payload
      return user
    }
    case actions.UPDATE_CONNECTED_USER_PROPERTY: {
      const { property, value } = action.payload
      return set(state, property, value)
    }
    default:
      return state
  }
}

const auth = combineReducers({
  user,
  isLoading,
  form,
})

export default auth
