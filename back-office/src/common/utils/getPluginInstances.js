//@flow
import { map, concat, chain, mapValues } from 'lodash'
import { values } from 'lodash/fp'
import type { PluginInstance } from '../../dashboard'

export const extractPluginInstances = (plugins: PluginInstance[]) => {
  if (!plugins.length) return plugins

  const subPlugins = chain(plugins)
    .map('props')
    .map(values)
    .flatten()
    .filter({ type: 'pluginList' })
    .map('value')
    .flatten()
    .value()

  const pluginsWithoutSubs = plugins.map(plugin => ({
    ...plugin,
    props: mapValues(
      plugin.props,
      prop =>
        prop.type === 'pluginList'
          ? { ...prop, value: map(prop.value, 'instanceId') }
          : prop,
    ),
  }))

  return concat(pluginsWithoutSubs, extractPluginInstances(subPlugins))
}

export default extractPluginInstances
