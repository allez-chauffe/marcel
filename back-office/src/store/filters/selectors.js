// @flow
import { createSelector } from 'reselect'
import { filter, chain, pick } from 'lodash'
import { update } from 'immutadot'
import { pluginsSelector } from '../../plugins/selectors'
import { selectedPluginSelector } from '../../dashboard'
import type { State } from '../types'

export const pluginFilterSelector = (state: State) => state.filters.plugins
export const propsFilterSelector = (state: State) => state.filters.props

export const filterByName = (filterString: string) => (
  item: string,
): boolean => {
  const regexPatern: string = chain(filterString)
    .split('')
    .without(' ')
    .join('.*')
    .value()

  const regexp = RegExp(`.*${regexPatern}.*`, 'i')
  return regexp.test(item)
}

export const filteredPluginsSeletor = createSelector(
  pluginsSelector,
  pluginFilterSelector,
  (plugins, filterString) => filter(plugins, filterByName(filterString)),
)

export const selectedPluginPropsFilteredSelector = createSelector(
  selectedPluginSelector,
  propsFilterSelector,
  (plugin, filterString) =>
    update(plugin, 'props', () => filterByName(filterString)),
)
