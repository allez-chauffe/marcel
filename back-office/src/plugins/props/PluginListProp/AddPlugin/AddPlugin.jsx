//@flow
import React from 'react'
import { without } from 'lodash'
import Dropdown from 'react-toolbox/lib/dropdown/Dropdown'
import type { Plugin } from '../../../../plugins'

class AddPlugin extends React.Component {
  props: {
    availablePlugins: Plugin[],
    plugins: Plugin[],
    addPlugin: Plugin => void,
  }

  addPlugin = (plugin: Plugin) => {
    if (plugin !== -1) this.props.addPlugin(plugin)
  }

  render() {
    const { plugins, availablePlugins } = this.props
    const source = [
      { value: -1, label: 'Ajouter un plugin' },
      ...without(availablePlugins, ...plugins).map(plugin => ({
        value: plugin,
        label: plugin.name,
      })),
    ]

    return <Dropdown source={source} value={-1} onChange={this.addPlugin} />
  }
}

export default AddPlugin
