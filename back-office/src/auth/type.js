//@flow

export type AuthState = {
  token: ?string,
}

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

export type AuthAction = LoginSuccessAction | LoginFailAction | LogoutAction
