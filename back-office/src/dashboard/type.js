//@flow
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
}

export type PluginInstanceMap = { [instanceId: string]: PluginInstance }
export type Dashboard = {
  id: string,
  name: string,
  description: string,
  rows: number,
  cols: number,
  ratio: number,
  stylesvar: {
    'primary-color': string,
    'secondary-color': string,
    'background-color': string,
    'font-family': string,
  },
  plugins: PluginInstanceMap,
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

export type ConfirmDashboardDeletionAction = {
  type: 'DASHBOARD/CONFIRM_DASHBOARD_DELETION',
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
    plugin: Plugin,
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

export type ToggleDisplayGridAction = {
  type: 'DASHBOARD/TOGGLE_DISPLAY_GRID',
}

// eslint-disable-next-line no-use-before-define
export type DashboardThunk = ((DashboardAction) => mixed, () => State) => void

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
  | ConfirmDashboardDeletionAction
  | CancelDashboardDeletionAction
  | DeleteDashboardAction
  | AddDashboardAction
  | ToggleDisplayGridAction

export type DashboardState = {
  selectedPlugin: string | null,
  selectedDashboard: string | null,
  deletingDashboard: string | null,
  displayGrid: boolean,
  loading: boolean,
  dashboards: DashboardMap,
}
