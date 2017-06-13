//@flow
import type { SelectPluginAction } from './type'

export const actions = {
  SELECT_PLUGIN: 'DASHBOARD/SELECT_PLUGIN',
}

export const selectPlugin = (plugin: string): SelectPluginAction => ({
  type: actions.SELECT_PLUGIN,
  payload: { elementName: plugin },
})
