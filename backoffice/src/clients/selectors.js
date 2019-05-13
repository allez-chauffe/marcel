import { createSelector } from 'reselect'
import { values } from 'lodash/fp'

export const clientsMapSelector = state => state.clients.clients

export const clientsSelector = createSelector(
  clientsMapSelector,
  values,
)

export const isClientLoadingSelector = (state, client) => !!state.clients.loading[client.id]

export const associatingClientSelector = state => state.clients.associating
