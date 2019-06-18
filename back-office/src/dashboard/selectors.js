import { createSelector } from 'reselect'
import { keyBy } from 'lodash/fp'
import { flow, map, update } from 'immutadot'
import { mapPluginInstancesToProps } from '../common/utils'

export const dashboardsSelector = state => state.dashboard.dashboards

export const pluginInstancesSelector = state => state.dashboard.pluginInstances

export const selectedDashboardNameSelector = state => {
  // WORKAOURND: the route params are not stored in redux...
  const urlSegments = window.location.href.split('/')
  return urlSegments[urlSegments.length - 1]
}

export const selectedPluginNameSelector = state => state.dashboard.selectedPlugin

export const deletingDashboardSelector = state => state.dashboard.deletingDashboard

export const selectedDashboardSelector = createSelector(
  dashboardsSelector,
  pluginInstancesSelector,
  selectedDashboardNameSelector,
  (dashboards, pluginInstances, selectedName) => {
    if (!selectedName || !dashboards[selectedName]) return null
    return flow(
      map('plugins', mapPluginInstancesToProps(pluginInstances)),
      update('plugins', keyBy('instanceId')),
    )(dashboards[selectedName])
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
