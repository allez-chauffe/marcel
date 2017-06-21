//@flow
import { keyBy, values } from 'lodash'
import { pick } from 'lodash/fp'
import { dashboardSelector } from './selectors'
import type { Plugin, Prop } from '../plugins'
import type {
  SelectPluginAction,
  PluginInstance,
  AddPluginAction,
  DeletePluginAction,
  SaveLayoutAction,
  UpdateConfigAction,
  Layout,
  ChangePropAction,
  DashboardThunk,
} from './type'

export const actions = {
  SELECT_PLUGIN: 'DASHBOARD/SELECT_PLUGIN',
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

export const addPlugin = (plugin: Plugin): AddPluginAction => ({
  type: actions.ADD_PLUGIN,
  payload: { plugin },
})

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
  dispatch({ type: actions.UPLOAD_STARTED })

  const { name, description, plugins } = dashboardSelector(getState())
  const requestBody = {
    name,
    description,
    plugins: values(plugins).map(
      pick(['elementName', 'instanceId', 'props', 'x', 'y', 'columns', 'rows']),
    ),
  }

  console.log('Upload to the server : ', JSON.stringify(requestBody, null, 2))

  dispatch({ type: actions.UPLOAD_SUCCESSED })
}

export const updateConfig = (property: string) => (
  value: string | number,
): UpdateConfigAction => ({
  type: actions.UPDATE_CONFIG,
  payload: { property, value },
})
