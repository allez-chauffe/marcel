//@flow
import type { Plugin, Prop } from '../plugins'

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
  plugins: { [instanceId: string]: ?PluginInstance },
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

export type DeletePluginAction = {
  type: 'DASHBOARD/DELETE_PLUGIN',
  payload: {
    plugin: Plugin,
  },
}

export type ChangePropAction = {
  type: 'DASHBOARD/CHANGE_PROP',
  payload: {
    instanceId: string,
    prop: Prop,
    value: mixed,
  },
}

export type DashboardAction =
  | SelectPluginAction
  | AddPluginAction
  | DeletePluginAction
  | ChangePropAction

export type DashboardState = {
  selectedPlugin: string | null,
  dashboard: Dashboard,
}
