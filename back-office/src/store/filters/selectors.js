// @flow
import { createSelector } from 'reselect'
import { filter, chain } from 'lodash'
import { pickBy } from 'immutadot'
import { pluginsSelector } from '../../plugins/selectors'
import { selectedPluginSelector } from '../../dashboard'
import type { PluginInstance } from '../../dashboard'
import type { State } from '../types'

export const pluginFilterSelector = (state: State) => state.filters.plugins
export const propsFilterSelector = (state: State) => state.filters.props

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
