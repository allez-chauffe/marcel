import { actions as loadActions } from '../store/loaders'
import { actions } from './actions'
import { set } from 'immutadot/core/set'
import { combineReducers } from 'redux'

const list = (state = [], action) => {
  switch (action.type) {
    case loadActions.LOAD_PLUGINS_SUCCESSED: {
      return action.payload.plugins
    }
    case actions.PLUGIN_UPDATE_SUCCESS: {
      const pluginIndex = state.findIndex(
        plugin => plugin.eltName === action.payload.plugin.eltName,
      )
      return set(state, `[${pluginIndex}]`, action.payload.plugin)
    }
    case actions.ADD_PLUGIN_SUCCESS: {
      return [...state, action.payload.plugin]
    }
    case actions.PLUGIN_DELETION_SUCCESS: {
      return state.filter(plugin => plugin.eltName !== action.payload.eltName)
    }
    default:
      return state
  }
}

const updating = (state = {}, action) => {
  switch (action.type) {
    case actions.UPDATE_PLUGIN_REQUESTED:
    case actions.PLUGIN_DELETION_REQUESTED: {
      return { ...state, [action.payload.eltName]: true }
    }
    case actions.UPDATE_PLUGIN_LOADED:
    case actions.PLUGIN_DELETION_LOADED: {
      return { ...state, [action.payload.eltName]: false }
    }
    default:
      return state
  }
}

const adding = (state = false, action) => {
  if (action.type === actions.ADD_PLUGIN_REQUESTED) return true
  if (action.type === actions.ADD_PLUGIN_LOADED) return false
  return state
}

export default combineReducers({ list, updating, adding })
