//@flow
import { createSelector } from 'reselect'
import { values } from 'lodash/fp'
import type { State } from '../store'

export const clientsMapSelector = (state: State) => state.clients.clients

export const clientsSelector = createSelector(clientsMapSelector, values)

export const isClientLoadingSelector = (state: State, client: Client) =>
  !!state.clients.loading[client.id]
