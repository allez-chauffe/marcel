export const userSelector = state => state.auth.user

export const isLoggedInSelector = state => !!state.auth.user

export const isLoading = state => state.auth.isLoading

export const loginSelector = state => state.auth.form.login

export const passwordSelector = state => state.auth.form.password
