//@flow
import type { State } from '../store'

export const tokenSelector = (state: State) => state.auth.token

export const isLoggedInSelector = (state: State) => !!state.auth.token
