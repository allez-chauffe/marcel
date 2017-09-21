// @flow
import { createSelector } from 'reselect'
import { filter, chain, groupBy } from 'lodash'
import { pickBy } from 'immutadot'
import { pluginsSelector } from '../../plugins/selectors'
import { selectedPluginSelector, selectedDashboardSelector } from '../../dashboard'
import { clientsSelector } from '../../clients'
import type { PluginInstance } from '../../dashboard'
import type { State } from '../types'

export const pluginFilterSelector = (state: State) => state.filters.plugins
export const propsFilterSelector = (state: State) => state.filters.props
export const clientsFilterSelector = (state: State) => state.filters.clients

export const filterByName = (filterString: string) => {
  const regexPatern: string = chain(filterString)
    .split('')
    .without(' ')
    .join('.*')
    .value()

  const regexp = RegExp(`.*${regexPatern}.*`, 'i')

  return <T: { name: string }>(item: T): boolean => regexp.test(item.name)
}

export const filteredPluginsSeletor = createSelector(
  pluginsSelector,
  pluginFilterSelector,
  (plugins, filterString) => filter(plugins, filterByName(filterString)),
)

export const selectedPluginPropsFilteredSelector = createSelector(
  selectedPluginSelector,
  propsFilterSelector,
  (plugin, filterString): ?PluginInstance =>
    plugin && pickBy(plugin, 'props', filterByName(filterString)),
)

export const filteredClientsSelector = createSelector(
  clientsSelector,
  clientsFilterSelector,
  (clients, filterString) => filter(clients, filterByName(filterString)),
)

export const partionedFilteredClientsSelector = createSelector(
  selectedDashboardSelector,
  filteredClientsSelector,
  (dashboard, clients) => {
    return groupBy(clients, client => {
      if (dashboard && dashboard.id === client.mediaID) return 'associated'
      if (client.isConnected) return 'connected'
      return 'disconnected'
    })
  },
)
