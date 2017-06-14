//@flow
import { createSelector } from 'reselect'
import { find } from 'lodash'
import type { State } from '../store'
import { pluginsSelector } from '../plugins/selectors'

export const selectedPluginNameSelector = (state: State) =>
  state.dashboard.selectedPlugin

export const selectedPluginSelector = createSelector(
  pluginsSelector,
  selectedPluginNameSelector,
  (plugins, elementName) => find(plugins, { elementName }),
)
