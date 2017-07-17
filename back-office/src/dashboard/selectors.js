//@flow
import { createSelector } from 'reselect'
import { mapValues, keyBy } from 'lodash'
import { mapValues as i_mapValues } from 'immutadot'
import type { State } from '../store'
import type { Dashboard, DashboardMap } from './type'

export const dashboardsSelector = (state: State): DashboardMap =>
  state.dashboard.dashboards

export const pluginInstancesSelector = (state: State) =>
  state.dashboard.pluginInstances

export const selectedDashboardNameSelector = (state: State) =>
  state.dashboard.selectedDashboard

export const selectedPluginNameSelector = (state: State) =>
  state.dashboard.selectedPlugin

export const deletingDashboardSelector = (state: State) =>
  state.dashboard.deletingDashboard

export const displayGridSelector = (state: State) => state.dashboard.displayGrid

export const selectedDashboardSelector = createSelector(
  dashboardsSelector,
  pluginInstancesSelector,
  selectedDashboardNameSelector,
  (dashboards, pluginInstances, selectedName) => {
    if (!selectedName || !dashboards[selectedName]) return null
    const dashboard = dashboards[selectedName]
    return {
      ...dashboard,
      plugins: mapValues(
        keyBy(dashboard.plugins),
        instanceId => pluginInstances[instanceId],
      ),
    }
  },
)

export const selectedPluginSelector = createSelector(
  pluginInstancesSelector,
  selectedPluginNameSelector,
  (pluginInstances, instanceId) => {
    if (!instanceId || pluginInstances[instanceId]) return null
    const pluginInstance = pluginInstances[instanceId]
    return {
      ...pluginInstance,
      props: concat(
        reject(pluginInstance.prop, { type: 'pluginList' }),
        filter(pluginInstance.prop, { type: 'pluginList' }).map(
          set('value', pluginInstances),
        ),
      ),
    }
    return i_mapValues(
      pluginInstance,
      'props',
      prop =>
        prop.type === 'pluginList'
          ? { ...prop, value: pluginInstances[prop.value] }
          : prop,
    )
  },
)
