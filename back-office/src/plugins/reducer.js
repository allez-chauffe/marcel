import { actions as loadActions } from '../store/loaders'
import { actions } from './actions'
import { set } from 'immutadot/core/set'

const intialState = { list: [], updating: null }

const reducer = (state = intialState, action) => {
  switch (action.type) {
    case loadActions.LOAD_PLUGINS_SUCCESSED: {
      return set(state, 'list', action.payload.plugins)
    }
    case actions.PLUGIN_UPDATE_SUCCESS: {
      const pluginIndex = state.list.findIndex(
        plugin => plugin.eltName === action.payload.plugin.eltName,
      )
      return set(state, `list[${pluginIndex}]`, action.payload.plugin)
    }
    case actions.UPDATE_PLUGIN_REQUESTED: {
      return set(state, 'updating', true)
    }
    case actions.UPDATE_PLUGIN_LOADED: {
      return set(state, 'updating', false)
    }
    default:
      return state
  }
}

export default reducer
