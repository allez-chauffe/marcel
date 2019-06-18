import { userBackend } from '../api'
import { toastr } from 'react-redux-toastr'

export const actions = {
  ADD_USER_SUCCESS: 'USER/ADD_USER_SUCCESS',
  UPDATE_CURRENT_USER_PROPERTY: 'USER/UPDATE_CURRENT_USER_PROPERTY',
  UPDATE_USER_SUCCESS: 'USER/UPDATE_USER_SUCCESS',
  RESET_CURRENT_USER: 'USER/RESET_CURRENT_USER',
  EDIT_USER: 'USER/EDIT_USER',
  DELETE_USER_SUCCESS: 'USER/DELETE_USER_SUCCESS',
}

export const addUserSuccess = user => ({
  type: actions.ADD_USER_SUCCESS,
  payload: {
    user: user,
  },
})

export const addUser = user => dispatch => {
  userBackend
    .addUser(user)
    .then(userAdded => dispatch(addUserSuccess(userAdded)))
    .then(() => dispatch(resetCurrentUser()))
    .catch(error => {
      toastr.error("Création d'un utilisateur", "Erreur durant la création de l'utilisateur")
      throw error
    })
}

export const updateCurrentUserProperty = (property, value) => ({
  type: actions.UPDATE_CURRENT_USER_PROPERTY,
  payload: { property, value },
})

export const updateUserSuccess = user => ({
  type: actions.UPDATE_USER_SUCCESS,
  payload: {
    user: user,
  },
})

export const updateUser = user => dispatch => {
  userBackend
    .updateUser(user)
    .then(() => dispatch(updateUserSuccess(user)))
    .then(() => dispatch(resetCurrentUser()))
    .catch(error => {
      toastr.error("Mise à jour d'un utilisateur", "Erreur durant la mise à jour de l'utilisateur")
      throw error
    })
}

export const resetCurrentUser = () => ({
  type: actions.RESET_CURRENT_USER,
})

export const editUser = user => ({
  type: actions.EDIT_USER,
  payload: {
    user: user,
  },
})

export const deleteUserSuccess = id => ({
  type: actions.DELETE_USER_SUCCESS,
  payload: {
    id: id,
  },
})

export const deleteUser = id => dispatch => {
  userBackend
    .deleteUser(id)
    .then(() => dispatch(deleteUserSuccess(id)))
    .catch(error => {
      toastr.error("Suppression d'un utilisateur", "Erreur durant la suppression de l'utilisateur")
      throw error
    })
}
