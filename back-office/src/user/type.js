// @flow
export type UserAction = LoadersAction

export type UserState = {
  users: null,
  currentUser: null,
}

export type User = {
  id?: number,
  displayName: string,
  login: string,
  role: 'admin' | 'user',
  password?: string,
  confirmPassword?: string,
}

export type AddUserAction = {
  type: 'USER/ADD_USER_SUCCESS',
  payload: { user: User },
}

export type EditUserAction = {
  type: 'USER/EDIT_USER',
  payload: { user: User },
}

export type UpdateUserAction = {
  type: 'USER/UPDATE_USER_SUCCESS',
  payload: { user: User },
}

export type UpdateCurrentUserPropertyAction = {
  type: 'USER/UPDATE_CURRENT_USER_PROPERTY',
  payload: { 
    property: string,
    value: string,
  },
}

export type ResetCurrentUserAction = {
  type: 'USER/RESET_CURRENT_USER',
}

export type DeleteUserAction = {
  type: 'USER/DELETE_USER_SUCCESS',
  payload: {
    id: string,
  }
}

