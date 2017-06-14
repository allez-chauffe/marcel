//@flow
import type { Plugin } from '../plugins'

export type PluginInstance = Plugin & {
  instanceId: string,
  x: number,
  y: number,
  columns: number,
  rows: number,
}

export type Dashboard = {
  name: string,
  description: string,
  plugins: PluginInstance[],
}

// Redux
export type SelectPluginAction = {
  type: 'DASHBOARD/SELECT_PLUGIN',
  payload: {
    instanceId: string,
  },
}

export type AddPluginAction = {
  type: 'DASHBOARD/ADD_PLUGIN',
  payload: {
    plugin: Plugin,
  },
}

export type DashboardAction = SelectPluginAction | AddPluginAction

export type DashboardState = {
  selectedPlugin: string | null,
  dashboard: Dashboard,
}
