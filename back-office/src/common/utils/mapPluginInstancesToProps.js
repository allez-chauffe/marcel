//@flow
import { pickBy, merge, omitBy, mapValues } from 'lodash'
import type { PluginInstanceMap, PluginInstance } from '../../dashboard'

const mapPluginInstancesToProps = (pluginInstances: PluginInstanceMap) => (
  instanceId: string,
): PluginInstance => {
  const getPluginInstance = mapPluginInstancesToProps(pluginInstances)
  const { props, ...otherAttributes } = pluginInstances[instanceId]
  return {
    ...otherAttributes,
    props: merge(
      omitBy(props, { type: 'pluginList' }),
      mapValues(pickBy(props, { type: 'pluginList' }), prop => ({
        ...prop,
        value: prop.value.map(getPluginInstance),
      })),
    ),
  }
}

export default mapPluginInstancesToProps
