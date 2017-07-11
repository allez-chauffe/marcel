/* eslint-disable no-use-before-define */
//@flow
import type { Dispatch } from 'redux'
import type { State, Config } from '../types'
import type { Plugin } from '../../plugins'
import type { Dashboard } from '../../dashboard/type'

export type LoadingThunkAction = (Dispatch<*>, () => State) => void

export type LoadConfigStartedAction = {
  type: 'LOADERS/LOAD_CONFIG_STARTED',
}

export type LoadPluginsStartedAction = {
  type: 'LOADERS/LOAD_PLUGINS_STARTED',
}

export type LoadDashboardsStartedAction = {
  type: 'LOADERS/LOAD_DASHBOARDS_STARTED',
}

export type LoadConfigSuccessedAction = {
  type: 'LOADERS/LOAD_CONFIG_SUCCESSED',
  payload: { config: Config },
}

export type LoadPluginsSuccessedAction = {
  type: 'LOADERS/LOAD_PLUGINS_SUCCESSED',
  payload: { plugins: Plugin[] },
}

export type LoadDashboardsSuccessedAction = {
  type: 'LOADERS/LOAD_DASHBOARDS_SUCCESSED',
  payload: { dashboards: Dashboard[] },
}

export type LoadConfigFailedAction = {
  type: 'LOADERS/LOAD_CONFIG_FAILED',
  payload: { error: mixed },
}

export type LoadPluginsFailedAction = {
  type: 'LOADERS/LOAD_PLUGINS_FAILED',
  payload: { error: mixed },
}

export type LoadDashboardsFailedAction = {
  type: 'LOADERS/LOAD_DASHBOARDS_FAILED',
  payload: { error: mixed },
}

export type LoadersAction =
  | LoadConfigStartedAction
  | LoadConfigSuccessedAction
  | LoadConfigFailedAction
  | LoadPluginsStartedAction
  | LoadPluginsSuccessedAction
  | LoadPluginsFailedAction
  | LoadDashboardsStartedAction
  | LoadDashboardsSuccessedAction
  | LoadDashboardsFailedAction

export type LoadersState = {
  config: boolean,
  dashboards: boolean,
  plugins: boolean,
}
