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

export const deletingDashboardSelector = (state: State) =>
  state.dashboard.deletingDashboard

export const displayGridSelector = (state: State) => state.dashboard.displayGrid

export const selectedDashboardSelector = createSelector(
  dashboardsSelector,
  selectedDashboardNameSelector,
  (dashboards, selectedName) =>
    selectedName ? dashboards[selectedName] : null,
)

const findPlugin = (dashboard: ?Dashboard, instanceId: string | null) =>
  dashboard ? (instanceId ? dashboard.plugins[instanceId] : null) : null

export const selectedPluginSelector = createSelector(
  selectedDashboardSelector,
  selectedPluginNameSelector,
  findPlugin,
)
