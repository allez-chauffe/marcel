//@flow
import React from 'react'
import { without } from 'lodash'
import { SortableContainer, arrayMove } from 'react-sortable-hoc'
import List from 'react-toolbox/lib/list/List'
import type { Plugin } from '../../type'
import SortablePlugin from './SortablePlugin'

import './PluginListProp.css'

const SortablePluginList = SortableContainer(({ plugins, onDelete }) =>
  <List>
    {plugins.map((plugin, index) =>
      <SortablePlugin
        key={plugin.elementName}
        index={index}
        plugin={plugin}
        onDelete={onDelete}
      />,
    )}
  </List>,
)

class PluginListProp extends React.Component {
  props: {
    plugins: Plugin[],
    name: string,
    value: Plugin[],
    onChange: (Plugin[]) => void,
  }

  onSortEnd = (swap: { oldIndex: number, newIndex: number }) => {
    const { oldIndex, newIndex } = swap
    this.props.onChange(arrayMove(this.props.value, oldIndex, newIndex))
  }

  onDelete = (plugin: Plugin) =>
    this.props.onChange(without(this.props.value, plugin))

  render() {
    const { value } = this.props
    return (
      <SortablePluginList
        onDelete={this.onDelete}
        helperClass="sortablePlugin"
        plugins={value}
        onSortEnd={this.onSortEnd}
        useDragHandle={true}
      />
    )
  }
}

export default PluginListProp
