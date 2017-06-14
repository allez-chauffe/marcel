//@flow
import type { Plugin } from '../plugins'

import type {
  LayoutItem as LayoutItemT,
  Layout as LayoutT,
} from 'react-grid-layout/build/utils.js.flow'

export type Layout = LayoutT
export type LayoutItem = LayoutItemT
export type LayoutMap = { [instanceId: string]: ?LayoutItem }

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

export type SaveLayoutAction = {
  type: 'DASHBOARD/SAVE_LAYOUT',
  payload: {
    layout: LayoutMap,
  },
}

export type DashboardAction =
  | SelectPluginAction
  | AddPluginAction
  | SaveLayoutAction

export type DashboardState = {
  selectedPlugin: string | null,
  dashboard: Dashboard,
}
