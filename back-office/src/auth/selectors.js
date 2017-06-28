//@flow
import type { State } from '../store'

export const tokenSelector = (state: State) => state.auth.token

export const isLoggedInSelector = (state: State) => !!state.auth.token

export const loginSelector = (state: State) => state.auth.form.login

export const passwordSelector = (state: State) => state.auth.form.password
