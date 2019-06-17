import React, { Component } from 'react'
import { SortableContainer, arrayMove } from 'react-sortable-hoc'
import List from 'react-toolbox/lib/list/List'
import SortablePlugin from './SortablePlugin'
import AddPlugin from './AddPlugin'

import './PluginListProp.css'

const SortablePluginList = SortableContainer(({ plugins, onDelete }) => (
  <List>
    {plugins.map((plugin, index) => (
      <SortablePlugin key={plugin.eltName} index={index} plugin={plugin} onDelete={onDelete} />
    ))}
  </List>
))

class PluginListProp extends Component {
  onSortEnd = swap => {
    const { oldIndex, newIndex } = swap
    this.props.reorderSubPlugins(arrayMove(this.props.value, oldIndex, newIndex))
  }

  addPlugin = plugin => this.props.addSubPlugin(this.props.name, plugin)

  deletePlugin = plugin => this.props.deletePlugin(plugin)

  render() {
    const { value } = this.props
    return (
      <div>
        <SortablePluginList
          onDelete={this.deletePlugin}
          helperClass="sortablePlugin"
          plugins={value}
          onSortEnd={this.onSortEnd}
          useDragHandle={true}
        />
        <AddPlugin addPlugin={this.addPlugin} plugins={value} />
      </div>
    )
  }
}

export default PluginListProp
