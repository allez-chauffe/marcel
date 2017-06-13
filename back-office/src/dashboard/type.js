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
  type: string,
  payload: {
    elementName: string,
  },
}

export type DashboardAction = SelectPluginAction

export type DashboardState = {
  selectedPlugin: string | null,
  dashboard: Dashboard,
}
