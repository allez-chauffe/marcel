//@flow
import React from 'react'
import type { Plugin } from '../../type'
import ListItem from 'react-toolbox/lib/list/ListItem'
import IconButton from 'react-toolbox/lib/button/IconButton'

class PluginListItem extends React.Component {
  props: { plugin: Plugin, addPlugin: Plugin => void }

  onClick = () => this.props.addPlugin(this.props.plugin)

  render() {
    const { plugin } = this.props
    const { name, eltName, icon } = plugin
    return (
      <ListItem
        caption={name}
        ripple={false}
        key={eltName}
        legend={eltName}
        leftIcon={icon}
        rightIcon={<IconButton icon="add" primary onClick={this.onClick} />}
      />
    )
  }
}

export default PluginListItem
