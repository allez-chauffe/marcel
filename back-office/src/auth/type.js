//@flow
/* eslint-disable no-use-before-define */
import type { Dispatch } from 'redux'
import type { State } from '../store'
import { User } from '../user'

// export type User = {
//   id: string,
//   displayName: string,
//   role: 'user' | 'admin',
// }

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


  export type UpdateConnectedUserAction = {
    type: 'AUTH/UPDATE_CONNECTED_USER_SUCCESS',
    payload: { user: User },
  }
  
  export type UpdateConnectedUserPropertyAction = {
    type: 'AUTH/UPDATE_CONNECTED_USER_PROPERTY',
    payload: { 
      property: string,
      value: string,
    },
  }