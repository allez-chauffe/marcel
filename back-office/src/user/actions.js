//@flow
import type {
  AddUserAction,
  UpdateCurrentUserPropertyAction,
  UpdateUserAction,
  ResetCurrentUserAction,
} from './type'

import { userBackend } from '../api'

export const actions = {
  ADD_USER_SUCCESS: "USER/ADD_USER_SUCCESS",
  UPDATE_CURRENT_USER_PROPERTY: "USER/UPDATE_CURRENT_USER_PROPERTY",
  UPDATE_USER_SUCCESS: "USER/UPDATE_USER_SUCCESS",
  RESET_CURRENT_USER: "USER/RESET_CURRENT_USER",
  EDIT_USER: "USER/EDIT_USER",
  DELETE_USER_SUCCESS: "USER/DELETE_USER_SUCCESS"
}

export const addUserSuccess = (user: User): AddUserAction => ({
  type: actions.ADD_USER_SUCCESS,
  payload: {
    user: user
  }
})

export const addUser = (user: User) => (dispatch) => {
  userBackend.addUser(user)
    .then((userAdded) => dispatch(addUserSuccess(userAdded)))
    .then(() => dispatch(resetCurrentUser()))
    .catch(error => {
      console.error(error)
    })
}

export const updateCurrentUserProperty = (property: string, value: string): UpdateCurrentUserPropertyAction => ({
  type: actions.UPDATE_CURRENT_USER_PROPERTY,
  payload: { property, value },
})

export const updateUserSuccess = (user: User): UpdateUserAction => ({
  type: actions.UPDATE_USER_SUCCESS,
  payload: {
    user: user
  }
})

export const updateUser = (user: User) => (dispatch) => {
  userBackend.updateUser(user)
    .then(() => dispatch(updateUserSuccess(user)))
    .then(() => dispatch(resetCurrentUser()))
    .catch(error => {
      console.error(error)
    })
}

export const resetCurrentUser = (): ResetCurrentUserAction => ({
  type: actions.RESET_CURRENT_USER,
})

export const editUser = (user: User): EditUserAction => ({
  type: actions.EDIT_USER,
  payload: {
    user: user,
  }
})

export const deleteUserSuccess = (id: string): DeleteUserAction => ({
  type: actions.DELETE_USER_SUCCESS,
  payload: {
    id: id
  }
})

export const deleteUser = (id: string) => (dispatch) => {
  userBackend.deleteUser(id)
    .then(() => dispatch(deleteUserSuccess(id)))
    .catch(error => {
      console.error(error)
    })
}