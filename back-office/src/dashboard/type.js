//@flow
/* eslint-disable no-use-before-define */
import type { Dispatch } from 'redux'
import type { Plugin, Prop } from '../plugins'
import type { State } from '../store'

import type {
  LayoutItem as LayoutItemT,
  Layout as LayoutT,
} from 'react-grid-layout/build/utils.js.flow'

export type Layout = LayoutT
export type LayoutItem = LayoutItemT
export type LayoutMap = { [instanceId: string]: ?LayoutItem }
export type DashboardMap = { [dashboardName: string]: ?Dashboard } //eslint-disable-line

export type PluginInstance = Plugin & {
  instanceId: string,
  x: number,
  y: number,
  cols: number,
  rows: number,
  parent?: { plugin: string, prop: string },
}

export type PluginInstanceMap = { [instanceId: string]: ?PluginInstance }
export type Dashboard = {
  id: string,
  name: string,
  description: string,
  rows: number,
  cols: number,
  screenRatio: number,
  displayGrid: boolean,
  isactive: boolean,
  screenshot: string,
  stylesvar: {
    'primary-color': string,
    'secondary-color': string,
    'background-color': string,
    'font-family': string,
  },
  plugins: string[],
}

// Redux
export type SelectPluginAction = {
  type: 'DASHBOARD/SELECT_PLUGIN',
  payload: {
    instanceId: string,
  },
}

export type SelectDashboardAction = {
  type: 'DASHBOARD/SELECT_DASHBOARD',
  payload: {
    dashboardId: string,
  },
}

export type UnselectDashboardAction = {
  type: 'DASHBOARD/UNSELECT_DASHBOARD',
}

export type RequireDashboardDeletionAction = {
  type: 'DASHBOARD/REQUIRE_DASHBOARD_DELETION',
  payload: { dashboardId: string },
}

export type DashboardDeletionThunk = () => (
  dispatch: Dispatch<DashboardDeletedAction | CancelDashboardDeletionAction>,
  getState: () => State,
) => void

export type DashboardDeletedAction = {
  type: 'DASHBOARD/DASHBOARD_DELETED',
}

export type CancelDashboardDeletionAction = {
  type: 'DASHBOARD/CANCEL_DASHBOARD_DELETION',
}

export type DeleteDashboardAction = {
  type: 'DASHBOARD/DELETE_DASHBOARD',
  payload: { dashboardId: string },
}

export type AddDashboardAction = {
  type: 'DASHBOARD/ADD_DASHBOARD',
  payload: { dashboard: Dashboard },
}

export type AddDashboardThunkAction = (Dispatch<AddDashboardAction>) => void

export type AddSubPluginAction = {
  type: 'DASHBOARD/ADD_SUB_PLUGIN',
  payload: {
    propName: string,
    plugin: Plugin,
  },
}

export type AddPluginAction = {
  type: 'DASHBOARD/ADD_PLUGIN',
  payload: {
    plugin: Plugin,
    x: number,
    y: number,
  },
}

export type AddPluginThunkAction = (
  dispatch: Dispatch<AddPluginAction>,
  getState: () => State,
) => void

export type DeletePluginAction = {
  type: 'DASHBOARD/DELETE_PLUGIN',
  payload: {
    plugin: PluginInstance,
  },
}
export type SaveLayoutAction = {
  type: 'DASHBOARD/SAVE_LAYOUT',
  payload: {
    layout: LayoutMap,
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

export type UploadStartedAction = {
  type: 'DASHBOARD/UPLOAD_STARTED',
}

export type UploadSuccesedAction = {
  type: 'DASHBOARD/UPLOAD_SUCCESSED',
}

export type UploadFailedAction = {
  type: 'DASHBOARD/UPLOAD_FAILED',
  payload: { error: string },
}

export type UpdateConfigAction = {
  type: 'DASHBOARD/UPDATE_CONFIG',
  payload: { property: string, value: string | number },
}

export type SelectPluginParentAction = {
  type: 'DASHBOARD/SELECT_PLUGIN_PARENT',
}

export type ReorderSubPluginAction = {
  type: 'DASHBOARD/REORDER_SUB_PLUGINS',
  payload: {
    instanceIds: string[],
    parent: {
      plugin: string,
      prop: string,
    },
  },
}

export type ActivateDashboardAction = {
  type: 'DASHBOARD/ACTIVATE_DASHBOARD',
  payload: {
    dashboardId: string,
  },
}

export type ActivateDashboardThunk = (Dispatch<ActivateDashboardAction>) => void

export type DeactivateDashboardAction = {
  type: 'DASHBOARD/DEACTIVATE_DASHBOARD',
  payload: {
    dashboardId: string,
  },
}

export type DeactivateDashboardThunk = (Dispatch<DeactivateDashboardAction>) => void

// eslint-disable-next-line no-use-before-define
export type DashboardThunk = (Dispatch<DashboardAction>, () => State) => void

export type DashboardAction =
  | SelectPluginAction
  | SelectDashboardAction
  | UnselectDashboardAction
  | AddPluginAction
  | DeletePluginAction
  | ChangePropAction
  | SaveLayoutAction
  | UploadStartedAction
  | UploadSuccesedAction
  | UploadFailedAction
  | RequireDashboardDeletionAction
  | DashboardDeletionThunk
  | DashboardDeletedAction
  | CancelDashboardDeletionAction
  | DeleteDashboardAction
  | AddDashboardAction
  | AddSubPluginAction

export type DashboardState = {
  selectedPlugin: string | null,
  selectedDashboard: string | null,
  deletingDashboard: string | null,
  loading: boolean,
  dashboards: DashboardMap,
  pluginInstances: PluginInstanceMap,
}
