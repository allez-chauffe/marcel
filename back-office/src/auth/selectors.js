//@flow
import type { State } from '../store'

export const userSelector = (state: State) => state.auth.user

export const isLoggedInSelector = (state: State) => !!state.auth.user

export const isLoading = (state: State) => state.auth.isLoading

export const loginSelector = (state: State) => state.auth.form.login

export const passwordSelector = (state: State) => state.auth.form.password
