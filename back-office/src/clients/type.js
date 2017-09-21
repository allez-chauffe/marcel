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

export type RequireClientAssociationAction = {
  type: 'CLIENTS/REQUIRE_CLIENT_ASSOCIATION',
  payload: {
    client: Client,
  },
}

export type ConfirmClientAssociationAction = {
  type: 'CLIENTS/CONFIRM_CLIENT_ASSOCIATION',
}

export type CancelClientAssociationAction = {
  type: 'CLIENTS/CANCEL_CLIENT_ASSOCIATION',
}

export type ClientState = {
  clients: ClientMap,
  loading: { [clientId: string]: boolean },
  associating: ?Client,
}

export type ClientAction =
  | LoadClientsStartedAction
  | LoadClientsSuccessedAction
  | LoadClientsFailedAction
  | AssociateClientThunk
  | RequireClientAssociationAction
  | ConfirmClientAssociationAction
  | CancelClientAssociationAction
