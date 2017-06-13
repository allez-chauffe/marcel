//@flow
import type { Plugin } from '../plugins'
import type { SelectPluginAction } from './type'

export const actions = {
  SELECT_PLUGIN: 'DASHBOARD/SELECT_PLUGIN',
}

export const selectPlugin = (plugin: Plugin): SelectPluginAction => ({
  type: actions.SELECT_PLUGIN,
  payload: { elementName: plugin.elementName },
})
