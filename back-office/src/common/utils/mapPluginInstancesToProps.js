import { map } from 'immutadot/array/map'
import { mapValues } from 'immutadot-lodash/object/mapValues'

const mapPluginInstancesToProps = pluginInstances => {
  const getPluginInstance = instanceId =>
    mapValues(pluginInstances[instanceId], 'props', prop =>
      prop.type === 'pluginList' ? map(prop, 'value', getPluginInstance) : prop,
    )
  return getPluginInstance
}

export default mapPluginInstancesToProps
