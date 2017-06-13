//@flow
import type { Plugin } from '../plugins'
import type {
  SelectPluginAction,
  PluginInstance,
  AddPluginAction,
} from './type'

export const actions = {
  SELECT_PLUGIN: 'DASHBOARD/SELECT_PLUGIN',
  ADD_PLUGIN: 'DASHBOARD/ADD_PLUGIN',
}

export const selectPlugin = (plugin: PluginInstance): SelectPluginAction => {
  console.log(plugin)
  return {
    type: actions.SELECT_PLUGIN,
    payload: { instanceId: plugin.instanceId },
  }
}

export const addPlugin = (plugin: Plugin): AddPluginAction => ({
  type: actions.ADD_PLUGIN,
  payload: { plugin },
})
