//@flow
import { createSelector } from 'reselect'
import { find, values } from 'lodash'
import type { State } from '../store'
import type { PluginInstance } from './type'

export const dashboardSelector = (state: State) => state.dashboard.dashboard

export const dashboardPluginsSelector = (state: State): PluginInstance[] =>
  values(state.dashboard.dashboard.plugins)

export const selectedPluginNameSelector = (state: State) =>
  state.dashboard.selectedPlugin

const findPlugin = (plugins: PluginInstance[], instanceId: ?string) => {
  return find(plugins, { instanceId })
}

export const selectedPluginSelector = createSelector(
  dashboardPluginsSelector,
  selectedPluginNameSelector,
  findPlugin,
)
