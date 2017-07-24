//@flow
import { pickBy, merge, omitBy } from 'lodash'
import { map, mapValues } from 'immutadot'
import type { PluginInstanceMap, PluginInstance } from '../../dashboard'

const mapPluginInstancesToProps = (pluginInstances: PluginInstanceMap) => (
  instanceId: string,
): PluginInstance => {
  const getPluginInstance = mapPluginInstancesToProps(pluginInstances)

  return mapValues(
    pluginInstances[instanceId],
    `props`,
    prop =>
      prop.type === 'pluginList' ? map(prop, `value`, getPluginInstance) : prop,
  )
}

export default mapPluginInstancesToProps
