import { map, concat, chain } from 'lodash'
import { mapValues } from 'immutadot'
import { values } from 'lodash/fp'

export const extractPluginInstances = plugins => {
  if (!plugins.length) return plugins

  const pluginsWithoutSubs = plugins.map(plugin =>
    mapValues(plugin, 'props', prop =>
      prop.type === 'pluginList' ? { ...prop, value: map(prop.value, 'instanceId') } : prop,
    ),
  )

  const propValues = chain(plugins)
    .map('props')
    .map(values)
    .flatten()
  const pluginListProps = propValues.filter({ type: 'pluginList' })
  const subPlugins = pluginListProps
    .map('value')
    .flatten()
    .value()

  return concat(pluginsWithoutSubs, extractPluginInstances(subPlugins))
}

export default extractPluginInstances
