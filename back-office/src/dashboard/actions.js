//@flow
import { keyBy, values, range, forEach, findIndex, some } from 'lodash'
import { pick } from 'lodash/fp'
import { toastr } from 'react-redux-toastr'

import { selectedDashboardSelector } from './selectors'
import type { Plugin, Prop } from '../plugins'
import type {
  SelectPluginAction,
  PluginInstance,
  AddPluginThunkAction,
  DeletePluginAction,
  SaveLayoutAction,
  UpdateConfigAction,
  Layout,
  Dashboard,
  ChangePropAction,
  DashboardThunk,
  SelectDashboardAction,
  UnselectDashboardAction,
  AddDashboardAction,
  DeleteDashboardAction,
  RequireDashboardDeletionAction,
  ConfirmDashboardDeletionAction,
  CancelDashboardDeletionAction,
} from './type'

export const actions = {
  SELECT_PLUGIN: 'DASHBOARD/SELECT_PLUGIN',
  SELECT_DASHBOARD: 'DASHBOARD/SELECT_DASHBOARD',
  UNSELECT_DASHBOARD: 'DASHBOARD/UNSELECT_DASHBOARD',
  REQUIRE_DASHBOARD_DELETION: 'DASHBOARD/REQUIRE_DASHBOARD_DELETION',
  CONFIRM_DASHBOARD_DELETION: 'DASHBOARD/CONFIRM_DASHBOARD_DELETION',
  CANCEL_DASHBOARD_DELETION: 'DASHBOARD/CANCEL_DASHBOARD_DELETION',
  DELETE_DASHBOARD: 'DASHBOARD/DELETE_DASHBOARD',
  ADD_DASHBOARD: 'DASHBOARD/ADD_DASHBOARD',
  ADD_PLUGIN: 'DASHBOARD/ADD_PLUGIN',
  DELETE_PLUGIN: 'DASHBOARD/DELETE_PLUGIN',
  CHANGE_PROP: 'DASHBOARD/CHANGE_PROP',
  SAVE_LAYOUT: 'DASHBOARD/SAVE_LAYOUT',
  UPLOAD_STARTED: 'DASHBOARD/UPLOAD_STARTED',
  UPLOAD_SUCCESSED: 'DASHBOARD/UPLOAD_SUCCESSED',
  UPLOAD_FAILED: 'DASHBOARD/UPLOAD_FAILED',
  UPDATE_CONFIG: 'DASHBOARD/UPDATE_CONFIG',
}

export const selectPlugin = (plugin: PluginInstance): SelectPluginAction => ({
  type: actions.SELECT_PLUGIN,
  payload: { instanceId: plugin.instanceId },
})

export const selectDashboard = (
  dashboard: Dashboard,
): SelectDashboardAction => ({
  type: actions.SELECT_DASHBOARD,
  payload: { dashboardId: dashboard.id },
})

export const unselectDashboard = (): UnselectDashboardAction => ({
  type: actions.UNSELECT_DASHBOARD,
})

export const requireDashboardDeletion = (
  dashboard: Dashboard,
): RequireDashboardDeletionAction => ({
  type: actions.REQUIRE_DASHBOARD_DELETION,
  payload: { dashboardId: dashboard.id },
})

export const confirmDashboardDeletion = (): ConfirmDashboardDeletionAction => ({
  type: actions.CONFIRM_DASHBOARD_DELETION,
})

export const cancelDashboardDeletion = (): CancelDashboardDeletionAction => ({
  type: actions.CANCEL_DASHBOARD_DELETION,
})

export const deleteDashboard = (
  dashboard: Dashboard,
): DeleteDashboardAction => ({
  type: actions.DELETE_DASHBOARD,
  payload: { dashboardId: dashboard.id },
})

export const addDashboard = (): AddDashboardAction => ({
  type: actions.ADD_DASHBOARD,
})

export const addPlugin = (plugin: Plugin): AddPluginThunkAction => (
  dispatch,
  getState,
) => {
  const dashboard: ?Dashboard = selectedDashboardSelector(getState())
  if (!dashboard)
    return toastr.error(
      "Erreur: Impossible d'ajouter un plugin",
      'Un dashboard doit être sélectionné',
    )

  const { rows, cols, plugins: pluginsMap } = dashboard
  const plugins: PluginInstance[] = values(pluginsMap)

  const freeSpaceMatrix = range(cols).map(() => range(rows).map(() => true))

  forEach(plugins, plugin =>
    range(plugin.x, plugin.x + plugin.columns).forEach(x =>
      range(plugin.y, plugin.y + plugin.rows).forEach(y => {
        freeSpaceMatrix[x][y] = false
      }),
    ),
  )

  const freeSpaceX = findIndex(freeSpaceMatrix, some)
  if (freeSpaceX === -1)
    return toastr.error('Erreur !', "Il n'y a plus de place pour ce plugin")
  const freeSpaceY = findIndex(freeSpaceMatrix[freeSpaceX])

  dispatch({
    type: actions.ADD_PLUGIN,
    payload: { plugin, x: freeSpaceX, y: freeSpaceY },
  })
}

export const deletePlugin = (plugin: Plugin): DeletePluginAction => ({
  type: actions.DELETE_PLUGIN,
  payload: { plugin },
})

export const changeProp = (
  plugin: PluginInstance,
  prop: Prop,
  value: mixed,
): ChangePropAction => ({
  type: actions.CHANGE_PROP,
  payload: {
    instanceId: plugin.instanceId,
    prop,
    value,
  },
})

export const saveLayout = (layout: Layout): SaveLayoutAction => ({
  type: actions.SAVE_LAYOUT,
  payload: { layout: keyBy(layout, 'i') },
})

export const uploadLayout = (): DashboardThunk => (dispatch, getState) => {
  const dashboard = selectedDashboardSelector(getState())
  if (!dashboard) throw new Error('A dashboard should be selected')

  dispatch({ type: actions.UPLOAD_STARTED })

  const { name, description, plugins } = dashboard
  const requestBody = {
    name,
    description,
    plugins: values(plugins).map(
      pick(['elementName', 'instanceId', 'props', 'x', 'y', 'columns', 'rows']),
    ),
  }

  // eslint-disable-next-line no-console
  console.log('Upload to the server : ', JSON.stringify(requestBody, null, 2))

  dispatch({ type: actions.UPLOAD_SUCCESSED })
  toastr.success('Enregistré', 'Le dashboard à bien été enregistré')
}

export const updateConfig = (property: string) => (
  value: string | number,
): UpdateConfigAction => ({
  type: actions.UPDATE_CONFIG,
  payload: { property, value },
})
