import { backend } from '../api'
import { toastr } from 'react-redux-toastr'

export const actions = {
  UPDATE_PLUGIN_REQUESTED: 'UPDATE_PLUGIN_REQUESTED',
  PLUGIN_UPDATE_SUCCESS: 'PLUGIN_UPDATE_SUCCESS',
  UPDATE_PLUGIN_LOADED: 'UPDATE_PLUGIN_LOADED',

  ADD_PLUGIN_REQUESTED: 'ADD_PLUGIN_REQUESTED',
  ADD_PLUGIN_SUCCESS: 'ADD_PLUGIN_SUCCESS',
  ADD_PLUGIN_LOADED: 'ADD_PLUGIN_LOADED',

  PLUGIN_DELETION_REQUESTED: 'PLUGIN_DELETION_REQUESTED',
  PLUGIN_DELETION_SUCCESS: 'PLUGIN_DELETION_SUCCESS',
  PLUGIN_DELETION_LOADED: 'PLUGIN_DELETION_LOADED',
}

const pluginUpdateRequested = eltName => ({
  type: actions.UPDATE_PLUGIN_REQUESTED,
  payload: { eltName },
})

const pluginUpdateSuccess = plugin => ({
  type: actions.PLUGIN_UPDATE_SUCCESS,
  payload: { plugin },
})

export const updatePlugin = pluginEltName => async dispatch => {
  dispatch(pluginUpdateRequested(pluginEltName))
  try {
    const updatedPlugin = await backend.updatePlugin(pluginEltName)
    dispatch(pluginUpdateSuccess(updatedPlugin))
  } catch (err) {
    console.error(err)
    toastr.error('Mise à jour du plugin', 'Erreur durant la mise à jour')
  } finally {
    dispatch({ type: actions.UPDATE_PLUGIN_LOADED })
  }
}

const addPluginSuccess = plugin => ({
  type: actions.ADD_PLUGIN_SUCCESS,
  payload: { plugin },
})

export const addPlugin = pluginUrl => async dispatch => {
  dispatch({ type: actions.ADD_PLUGIN_REQUESTED })
  try {
    const addedPlugin = await backend.addPlugin(pluginUrl)
    dispatch(addPluginSuccess(addedPlugin))
  } catch (err) {
    console.error(err)
    let message = "Erreur durant l'ajout du plugin"
    if (err.text) message += `: ${await err.text()}`
    toastr.error('Ajout du plugin', message)
  } finally {
    dispatch({ type: actions.ADD_PLUGIN_LOADED })
  }
}

const pluginDeletionRequested = eltName => ({
  type: actions.PLUGIN_DELETION_REQUESTED,
  payload: { eltName },
})

const pluginDeletionSuccess = eltName => ({
  type: actions.PLUGIN_DELETION_SUCCESS,
  payload: { eltName },
})

export const deletePlugin = pluginEltName => async dispatch => {
  dispatch(pluginDeletionRequested(pluginEltName))

  try {
    await backend.deletePlugin(pluginEltName)
    dispatch(pluginDeletionSuccess(pluginEltName))
  } catch (err) {
    console.error(err)
    toastr.error('Suppression du plugin', 'Erreur lors de la suppression')
  } finally {
    dispatch({ type: actions.PLUGIN_DELETION_LOADED })
  }
}
