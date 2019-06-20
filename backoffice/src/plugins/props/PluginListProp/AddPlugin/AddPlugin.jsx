import React from 'react'
import { without } from 'lodash'
import Dropdown from 'react-toolbox/lib/dropdown/Dropdown'

class AddPlugin extends React.Component {
  addPlugin = plugin => {
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
