//@flow
import { createSelector } from 'reselect'
import type { State } from '../store'
import type { Dashboard, DashboardMap } from './type'

export const dashboardsSelector = (state: State): DashboardMap =>
  state.dashboard.dashboards

export const selectedDashboardNameSelector = (state: State) =>
  state.dashboard.selectedDashboard

export const selectedPluginNameSelector = (state: State) =>
  state.dashboard.selectedPlugin

export const selectedDashboardSelector = createSelector(
  dashboardsSelector,
  selectedDashboardNameSelector,
  (dashboards, selectedName) =>
    selectedName ? dashboards[selectedName] : null,
)

const findPlugin = (dashboard: ?Dashboard, instanceId: string | null) => {
  console.log({ dashboard, instanceId, plugin: dashboard.plugins[instanceId] })
  return dashboard ? (instanceId ? dashboard.plugins[instanceId] : null) : null
}

export const selectedPluginSelector = createSelector(
  selectedDashboardSelector,
  selectedPluginNameSelector,
  findPlugin,
)
