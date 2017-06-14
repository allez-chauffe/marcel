//@flow
import { keyBy } from 'lodash'
import type { Plugin } from '../plugins'
import type {
  SelectPluginAction,
  PluginInstance,
  AddPluginAction,
  SaveLayoutAction,
  Layout,
} from './type'

export const actions = {
  SELECT_PLUGIN: 'DASHBOARD/SELECT_PLUGIN',
  ADD_PLUGIN: 'DASHBOARD/ADD_PLUGIN',
  SAVE_LAYOUT: 'DASHBOARD/SAVE_LAYOUT',
}

export const selectPlugin = (plugin: PluginInstance): SelectPluginAction => ({
  type: actions.SELECT_PLUGIN,
  payload: { instanceId: plugin.instanceId },
})

export const addPlugin = (plugin: Plugin): AddPluginAction => ({
  type: actions.ADD_PLUGIN,
  payload: { plugin },
})

export const saveLayout = (layout: Layout): SaveLayoutAction => ({
  type: actions.SAVE_LAYOUT,
  payload: { layout: keyBy(layout, 'i') },
})
