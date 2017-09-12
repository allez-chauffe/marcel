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

export type PropTypes = {
  plugin: PluginInstance,
  onDelete: () => void,
  onSelect: () => void,
}

const SortablePlugin = (props: PropTypes) => {
  const { onSelect, onDelete, plugin } = props
  const iconStyle = { cursor: 'pointer' }
  return (
    <ListItem
      caption={plugin.name}
      ripple={false}
      leftIcon={<DragHandle />}
      rightIcon={
        <div>
          <FontIcon value="edit" onClick={onSelect} style={iconStyle} />
          <FontIcon value="delete" onClick={onDelete} style={iconStyle} />
        </div>
      }
    />
  )
}

export default SortablePlugin
