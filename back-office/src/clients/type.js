//@flow
/*eslint-disable no-use-before-define*/
import type { Dispatch } from 'redux'
import type { State } from '../store'
import type {
  LoadClientsStartedAction,
  LoadClientsSuccessedAction,
  LoadClientsFailedAction,
} from '../store/loaders'

export type Client = {
  id: string,
  name: string,
  type: string,
  mediaID: string,
}

export type ClientMap = { [id: string]: ?Client }

export type AssociateClientThunk = (Dispatch<ClientAssociationAction>, () => State) => void

export type ClientAssociationStartedAction = {
  type: 'CLIENTS/CLIENT_ASSOCIATION_STARTED',
  payload: {
    client: Client,
  },
}

export type ClientAssociationSuccessAction = {
  type: 'CLIENTS/CLIENT_ASSOCIATION_SUCCESS',
  payload: { client: Client },
}

export type ClientAssociationFailedAction = {
  type: 'CLIENTS/CLIENT_ASSOCIATION_FAILED',
  payload: { client: Client },
}

export type ClientAssociationAction =
  | ClientAssociationStartedAction
  | ClientAssociationSuccessAction
  | ClientAssociationFailedAction

export type ClientState = {
  clients: ClientMap,
  loading: { [clientId: string]: boolean },
}

export type ClientAction =
  | LoadClientsStartedAction
  | LoadClientsSuccessedAction
  | LoadClientsFailedAction
  | AssociateClientThunk
