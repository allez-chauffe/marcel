import { createSelector } from 'reselect'
import { keyBy } from 'lodash/fp'
import { chain } from 'immutadot'
import { mapPluginInstancesToProps } from '../common/utils'

export const dashboardsSelector = state => state.dashboard.dashboards

export const pluginInstancesSelector = state => state.dashboard.pluginInstances

export const selectedDashboardNameSelector = state =>
  state.router.params && state.router.params.mediaID

export const selectedPluginNameSelector = state => state.dashboard.selectedPlugin

export const deletingDashboardSelector = state => state.dashboard.deletingDashboard

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
