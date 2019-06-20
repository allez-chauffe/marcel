import { backend } from '../api'
import { toastr } from 'react-redux-toastr'

export const actions = {
  PLUGIN_UPDATE_SUCCESS: 'PLUGIN_UPDATE_SUCCESS',
  UPDATE_PLUGIN_REQUESTED: 'UPDATE_PLUGIN_REQUESTED',
  UPDATE_PLUGIN_LOADED: 'UPDATE_PLUGIN_LOADED',
}

const pluginUpdateSuccess = plugin => ({
  type: actions.PLUGIN_UPDATE_SUCCESS,
  payload: { plugin },
})

export const updatePlugin = pluginEltName => async dispatch => {
  dispatch({ type: actions.UPDATE_PLUGIN_REQUESTED })
  try {
    const updatedPlugin = await backend.updatePlugin(pluginEltName)
    dispatch(pluginUpdateSuccess(updatedPlugin))
  } catch (err) {
    toastr.error('Mise à jour du plugin', 'Erreur durant la mise à jour')
  } finally {
    dispatch({ type: actions.UPDATE_PLUGIN_LOADED })
  }
}
