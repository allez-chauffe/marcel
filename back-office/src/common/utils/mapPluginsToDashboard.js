//@flow
import { omit, find, mapValues } from 'lodash'
import type { Plugin } from '../../plugins'
import type { Dashboard } from '../../dashboard/type'

const mapPluginsToDashboard = (pluginCatalog: Plugin[]) => (
  dashboard: Dashboard,
) => {
  const plugins = mapValues(dashboard.plugins, plugin => {
    const pluginInstance = omit(plugin.frontend, ['files'])
    const { eltName } = pluginInstance
    const { instanceId } = plugin
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
    plugins,
  }
}

export default mapPluginsToDashboard
