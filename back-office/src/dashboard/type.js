//@flow
import type { Plugin, Prop } from '../plugins'

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

export type ChangePropAction = {
  type: 'DASHBOARD/CHANGE_PROP',
  payload: {
    instanceId: string,
    prop: Prop,
    value: mixed,
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
  | ChangePropAction

export type DashboardState = {
  selectedPlugin: string | null,
  dashboard: Dashboard,
}
