//@flow
/* eslint-disable no-use-before-define */
import type { Dispatch } from 'redux'
import type { State } from '../store'

export type User = {
  id: string,
  displayName: string,
  role: 'user' | 'admin',
}

export type AuthState = {
  user: ?User,
  isLoading: boolean,
  form: {
    login: string,
    password: string,
  },
}

export type LoginAction = () => (
  dispatch: Dispatch<LoginRequest | LoginSuccessAction | LoginFailAction>,
  getState: () => State,
) => mixed

export type LoginRequest = {
  type: 'AUTH/LOGIN_REQUEST',
}

export type LoginSuccessAction = {
  type: 'AUTH/LOGIN_SUCCESS',
  payload: {
    user: User,
  },
}

export type LoginFailAction = {
  type: 'AUTH/LOGIN_FAIL',
}

export type LogoutAction = () => (
  dispatch: Dispatch<LogoutRequest | LogoutSuccessAction | LogoutFailAction>,
  getState: () => State,
) => mixed

export type RefreshLoginAction = () => (
  dispatch: Dispatch<LoginSuccessAction | LoginFailAction | LoginRequest>,
  getState: () => State,
) => mixed

export type LogoutRequest = {
  type: 'AUTH/LOGOUT_REQUEST',
}

export type LogoutSuccessAction = {
  type: 'AUTH/LOGOUT_SUCCESS',
}

export type LogoutFailAction = {
  type: 'AUTH/LOGOUT_FAIL',
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
