//@flow
import { pickBy, merge, omitBy } from 'lodash'
import type { PluginInstanceMap, PluginInstance } from '../../dashboard'

const mapPluginInstancesToProps = (pluginInstances: PluginInstanceMap) => (
  instanceId: string,
): PluginInstance => {
  const getPluginInstance = mapPluginInstancesToProps(pluginInstances)
  const { props, ...otherAttributes } = pluginInstances[instanceId]
  return {
    ...otherAttributes,
    props: merge(
      pickBy(props, { type: 'pluginList' }),
      omitBy(props, { type: 'pluginList' }).mapValues(prop => ({
        ...prop,
        value: prop.value.map(getPluginInstance),
      })),
    ),
  }
}

export default mapPluginInstancesToProps
