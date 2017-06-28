//@flow
import type { Dispatch } from 'redux'
import type { State } from '../store'

export type AuthState = {
  token: ?string,
  form: {
    login: string,
    password: string,
  },
}

export type LoginAction = (
  dispatch: Dispatch<*>,
  getState: () => State,
) => mixed

export type LoginSuccessAction = {
  type: 'AUTH/LOGIN_SUCCESS',
  payload: { token: string },
}

export type LoginFailAction = {
  type: 'AUTH/LOGIN_FAIL',
  payload: { error: string },
}

export type LogoutAction = {
  type: 'AUTH/LOGOUT',
}

export type ChangeLoginAction = {
  type: 'AUTH/CHANGE_LOGIN',
  payload: { login: string },
}

export type ChangePasswordAction = {
  type: 'AUTH/CHANGE_PASSWORD',
  payload: { password: string },
}

export type ResetFormAction = {
  type: 'AUTH/RESET_FORM',
}

export type AuthAction =
  | LoginSuccessAction
  | LoginFailAction
  | LogoutAction
  | ChangeLoginAction
  | ChangePasswordAction
  | ResetFormAction
