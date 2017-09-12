//@flow
import { omit, find, map, mapValues, keyBy } from 'lodash'
import type { Plugin } from '../../plugins'
import type { Dashboard } from '../../dashboard/type'

const mapPluginsToDashboard = (pluginCatalog: Plugin[]) => (dashboard: Dashboard) => {
  const plugins = map(dashboard.plugins, plugin => {
    const pluginInstance = omit(plugin.frontend, ['files'])
    const { eltName, instanceId } = plugin
    const pluginBase = find(pluginCatalog, { eltName })

    if (!pluginBase) throw new Error(`Not found : ${plugin.name} (${eltName})`)

    return {
      ...pluginBase,
      ...pluginInstance,
      instanceId,
      props: mapValues(pluginBase.props, prop => ({
        ...prop,
        value: pluginInstance.props[prop.name],
      })),
    }
  })

  return {
    ...dashboard,
    plugins: keyBy(plugins, 'instanceId'),
  }
}

export default mapPluginsToDashboard
