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
    const { name, elementName, icon } = plugin
    return (
      <ListItem
        caption={name}
        ripple={false}
        key={elementName}
        legend={elementName}
        leftIcon={icon}
        rightIcon={<IconButton icon="add" primary onClick={this.onClick} />}
      />
    )
  }
}

export default PluginListItem
