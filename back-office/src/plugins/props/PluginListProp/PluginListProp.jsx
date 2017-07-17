//@flow
import React from 'react'
import { without } from 'lodash'
import { SortableContainer, arrayMove } from 'react-sortable-hoc'
import List from 'react-toolbox/lib/list/List'
import type { Plugin } from '../../type'
import SortablePlugin from './SortablePlugin'
import AddPlugin from './AddPlugin'

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

  addPlugin = (plugin: Plugin) =>
    this.props.addSubPlugin(this.props.name, plugin)

  deletePlugin = (plugin: PluginInstance) =>
    this.props.deleteSubPlugin(this.props.name, plugin)

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
