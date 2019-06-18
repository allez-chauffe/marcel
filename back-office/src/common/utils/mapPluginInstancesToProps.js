import { map, mapValues } from 'immutadot'

const mapPluginInstancesToProps = pluginInstances => instanceId => {
  const getPluginInstance = mapPluginInstancesToProps(pluginInstances)

  return mapValues(pluginInstances[instanceId], `props`, prop =>
    prop.type === 'pluginList' ? map(prop, `value`, getPluginInstance) : prop,
  )
}

export default mapPluginInstancesToProps
