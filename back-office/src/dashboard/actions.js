//@flow
import type { Plugin, Prop } from '../plugins'
import type {
  SelectPluginAction,
  PluginInstance,
  AddPluginAction,
  ChangePropAction,
  DeletePluginAction,
} from './type'

export const actions = {
  SELECT_PLUGIN: 'DASHBOARD/SELECT_PLUGIN',
  ADD_PLUGIN: 'DASHBOARD/ADD_PLUGIN',
  CHANGE_PROP: 'DASHBOARD/CHANGE_PROP',
  DELETE_PLUGIN: 'DASHBOARD/DELETE_PLUGIN',
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
