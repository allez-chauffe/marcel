//@flow
import { createSelector } from 'reselect'
import { find } from 'lodash'
import type { State } from '../store'

export const dashboardSelector = (state: State) => state.dashboard.dashboard

export const dashboardPluginsSelector = (state: State) =>
  state.dashboard.dashboard.plugins

export const selectedPluginNameSelector = (state: State) =>
  state.dashboard.selectedPlugin

export const selectedPluginSelector = createSelector(
  dashboardPluginsSelector,
  selectedPluginNameSelector,
  (plugins, instanceId) => find(plugins, { instanceId }),
)
