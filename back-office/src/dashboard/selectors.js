//@flow
import { createSelector } from 'reselect'
import { keyBy } from 'lodash/fp'
import { chain } from 'immutadot'
import type { State } from '../store'
import type { DashboardMap } from './type'
import { mapPluginInstancesToProps } from '../common/utils'

export const dashboardsSelector = (state: State): DashboardMap => state.dashboard.dashboards

export const pluginInstancesSelector = (state: State) => state.dashboard.pluginInstances

export const selectedDashboardNameSelector = (state: State) => state.dashboard.selectedDashboard

export const selectedPluginNameSelector = (state: State) => state.dashboard.selectedPlugin

export const deletingDashboardSelector = (state: State) => state.dashboard.deletingDashboard

export const displayGridSelector = (state: State) => state.dashboard.displayGrid

export const selectedDashboardSelector = createSelector(
  dashboardsSelector,
  pluginInstancesSelector,
  selectedDashboardNameSelector,
  (dashboards, pluginInstances, selectedName) => {
    if (!selectedName || !dashboards[selectedName]) return null
    return chain(dashboards[selectedName])
      .map('plugins', mapPluginInstancesToProps(pluginInstances))
      .update('plugins', keyBy('instanceId'))
      .value()
  },
)

export const selectedPluginSelector = createSelector(
  pluginInstancesSelector,
  selectedPluginNameSelector,
  (pluginInstances, instanceId) => {
    if (!instanceId || !pluginInstances[instanceId]) return null
    return mapPluginInstancesToProps(pluginInstances)(instanceId)
  },
)
