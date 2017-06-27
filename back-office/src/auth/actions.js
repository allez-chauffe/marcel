//@flow
import type { Dispatch } from 'redux'

export const actions = {
  LOGIN_SUCCESS: 'AUTH/LOGIN_SUCCESS',
  LOGIN_FAIL: 'AUTH/LOGIN_FAIL',
  LOGOUT: 'AUTH/LOGOUT',
}

export const login = (login: string, password: string) => (
  dispatch: Dispatch<*>,
) => {
  dispatch({ type: actions.LOGIN_SUCCESS, payload: { token: 'myAuthToken' } })
}

export const logout = () => ({
  type: actions.LOGOUT,
})
