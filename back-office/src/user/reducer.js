//@flow
import type { Reducer } from 'redux'
import type { UserAction, UserState } from './type'
import { actions as loadActions } from '../store/loaders'
import { actions as userActions } from './actions'
import { set, unset } from 'immutadot'
import { keyBy } from 'lodash'

const initialState = {
  users: {},
  currentUser: {
    displayName: '',
    login: '',
    role: '',
    password: '',
    confirmPassword: '',
  },
}

const users: Reducer<UserState, UserAction> = (state = initialState, action) => {
  switch (action.type) {
    case loadActions.LOAD_USERS_SUCCESSED: {
      const { users } = action.payload
      return {
        ...state,
        users: keyBy(users, 'id'),
      }
    }
    case userActions.ADD_USER_SUCCESS: {
      const { user } = action.payload
      return set(state, `users.${user.id}`, user)
    }
    case userActions.UPDATE_USER_SUCCESS: {
      const { user } = action.payload
      return set(state, `users.${user.id}`, user)
    }
    case userActions.DELETE_USER_SUCCESS: {
      const { id } = action.payload
      return unset(state, `users.${id}`)
    }
    case userActions.EDIT_USER: {
      const { user } = action.payload
      user.confirmPassword = ''
      return set(state, 'currentUser', user)
    }
    case userActions.UPDATE_CURRENT_USER_PROPERTY: {
      const { property, value } = action.payload
      return set(state, `currentUser.${property}`, value)
    }
    case userActions.RESET_CURRENT_USER: {
      return set(state, 'currentUser', initialState.currentUser)
    }
    default:
      return state
  }
}

export default users
