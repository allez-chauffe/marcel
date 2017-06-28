//@flow
import type {
  LoginAction,
  LogoutAction,
  ChangeLoginAction,
  ChangePasswordAction,
  ResetFormAction,
} from './type'

import { loginSelector, passwordSelector } from './selectors'

export const actions = {
  LOGIN_SUCCESS: 'AUTH/LOGIN_SUCCESS',
  LOGIN_FAIL: 'AUTH/LOGIN_FAIL',
  LOGOUT: 'AUTH/LOGOUT',
  CHANGE_LOGIN: 'AUTH/CHANGE_LOGIN',
  CHANGE_PASSWORD: 'AUTH/CHANGE_PASSWORD',
  RESET_FORM: 'AUTH/RESET_FORM',
}

export const login = (): LoginAction => (dispatch, getState) => {
  const state = getState()
  const login = loginSelector(state)
  const password = passwordSelector(state)

  dispatch(
    login === 'admin' && password === 'admin'
      ? { type: actions.LOGIN_SUCCESS, payload: { token: 'myAuthToken' } }
      : {
          type: actions.LOGIN_FAIL,
          payload: { error: 'Wrong login or password' },
        },
  )
}

export const logout = (): LogoutAction => ({
  type: actions.LOGOUT,
})

export const changeLogin = (login: string): ChangeLoginAction => ({
  type: actions.CHANGE_LOGIN,
  payload: { login },
})

export const changePassword = (password: string): ChangePasswordAction => ({
  type: actions.CHANGE_PASSWORD,
  payload: { password },
})

export const resetForm = (): ResetFormAction => ({
  type: actions.RESET_FORM,
})
