//@flow
import React from 'react'
import { SortableHandle } from 'react-sortable-hoc'
import ListItem from 'react-toolbox/lib/list/ListItem'
import FontIcon from 'react-toolbox/lib/font_icon/FontIcon'
import type { PluginInstance } from '../../../../dashboard'

import './SortablePlugin.css'

const DragHandle = SortableHandle(() =>
  <FontIcon value="menu" className="grab" />,
)

class SortablePlugin extends React.Component {
  props: {
    plugin: PluginInstance,
    onDelete: PluginInstance => void,
    selectPlugin: PluginInstance => void,
  }

  onDelete = () => {
    this.props.onDelete(this.props.plugin)
  }

  onSelect = () => {
    this.props.selectPlugin(this.props.plugin)
  }

  render() {
    const iconStyle = { cursor: 'pointer' }
    return (
      <ListItem
        caption={this.props.plugin.name}
        ripple={false}
        leftIcon={<DragHandle />}
        rightIcon={
          <div>
            <FontIcon value="edit" onClick={this.onSelect} style={iconStyle} />
            <FontIcon
              value="delete"
              onClick={this.onDelete}
              style={iconStyle}
            />
          </div>
        }
      />
    )
  }
}

export default SortablePlugin
