// @flow
import { createSelector } from 'reselect'
import { filter, chain } from 'lodash'
import { pluginsSelector } from '../../plugins/selectors'
import type { State } from '../types'

export const pluginFilterSelector = (state: State) => state.filters.plugins
export const propsFilterSelector = (state: State) => state.filters.props

export const filteredPluginsSeletor = createSelector(
  pluginsSelector,
  pluginFilterSelector,
  (plugins, pluginFilter) => {
    const regexPatern: string = chain(pluginFilter)
      .split('')
      .without(' ')
      .join('.*')
      .value()

    const regexp = RegExp(`.*${regexPatern}.*`, 'i')
    return filter(plugins, ({ name }) => regexp.test(name))
  },
)
