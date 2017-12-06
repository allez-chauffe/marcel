//@flow
import { toastr } from 'react-redux-toastr'
import type {
  LoginAction,
  LogoutAction,
  ChangeLoginAction,
  ChangePasswordAction,
  ResetFormAction,
  RefreshLoginAction,
} from './type'
import { authBackend } from '../api'
import { loginSelector, passwordSelector } from './selectors'
import type { User } from './type'

export const actions = {
  LOGIN_REQUEST: 'AUTH/LOGIN_REQUEST',
  LOGIN_SUCCESS: 'AUTH/LOGIN_SUCCESS',
  LOGIN_FAIL: 'AUTH/LOGIN_FAIL',
  LOGOUT_REQUEST: 'AUTH/LOGOUT_REQUEST',
  LOGOUT_SUCCESS: 'AUTH/LOGOUT_SUCCESS',
  LOGOUT_FAIL: 'AUTH/LOGOUT_FAIL',
  CHANGE_LOGIN: 'AUTH/CHANGE_LOGIN',
  CHANGE_PASSWORD: 'AUTH/CHANGE_PASSWORD',
  RESET_FORM: 'AUTH/RESET_FORM',
  DISCONNECTED: 'AUTH/DISCONNECTED',
}

const handleLogin = (dispatch, promise) =>
  promise
    .then((user: User) => dispatch({ type: actions.LOGIN_SUCCESS, payload: { user } }))
    .catch(response => {
      if (response.status !== 403)
        toastr.error('Erreur', "Impossible de contacter le serveur d'authentification")
      dispatch({ type: actions.LOGIN_FAIL })
      throw response
    })

export const login: LoginAction = () => (dispatch, getState) => {
  dispatch({ type: actions.LOGIN_REQUEST })

  const state = getState()
  const login = loginSelector(state)
  const password = passwordSelector(state)

  handleLogin(dispatch, authBackend.login(login, password)).catch(({ status }) => {
    if (status === 403) toastr.error('Erreur', 'Login ou mot de passe incorrect')
  })
}

export const refreshLogin: RefreshLoginAction = () => (dispatch, getState) => {
  dispatch({ type: actions.LOGIN_REQUEST })
  handleLogin(dispatch, authBackend.login())
}

export const logout: LogoutAction = () => dispatch => {
  dispatch({ type: actions.LOGOUT_REQUEST })
  dispatch({ type: actions.LOGOUT_SUCCESS })
}

export const disconnected = () => ({
  type: actions.DISCONNECTED,
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
