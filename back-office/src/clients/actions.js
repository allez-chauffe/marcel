//@flow
import { selectedDashboardSelector } from '../dashboard'
import { toastr } from 'react-redux-toastr'
import { backend } from '../api'
import type { AssociateClientThunk, Client } from './type'

export const actions = {
  CLIENT_ASSOCIATION_STARTED: 'CLIENTS/CLIENT_ASSOCIATION_STARTED',
  CLIENT_ASSOCIATION_SUCCESS: 'CLIENTS/CLIENT_ASSOCIATION_SUCCESS',
  CLIENT_ASSOCIATION_FAILED: 'CLIENTS/CLIENT_ASSOCIATION_FAILED',
}

export const associateClient = (client: Client): AssociateClientThunk => (dispatch, getState) => {
  dispatch({ type: actions.CLIENT_ASSOCIATION_STARTED, payload: { client } })

  const media = selectedDashboardSelector(getState())
  if (!media) {
    dispatch({ type: actions.CLIENT_ASSOCIATION_FAILED, payload: { client } })
    toastr.error("Impossible d'associer le client", "Aucun media n'est sélectioné")
    return
  }

  const newClient = { ...client, mediaID: media.id }

  backend
    .updateClient(newClient)
    .then(() =>
      dispatch({ type: actions.CLIENT_ASSOCIATION_SUCCESS, payload: { client: newClient } }),
    )
    .catch(error => {
      toastr.error("Erreur lors de l'association du client", error)
      dispatch({ type: actions.CLIENT_ASSOCIATION_FAILED, payload: { client: newClient } })
    })
}
