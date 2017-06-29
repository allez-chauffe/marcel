//@flow
import React from 'react'
import {
  SortableElement,
  SortableContainer,
  SortableHandle,
  arrayMove,
} from 'react-sortable-hoc'
import List from 'react-toolbox/lib/list/List'
import ListItem from 'react-toolbox/lib/list/ListItem'
import FontIcon from 'react-toolbox/lib/font_icon/FontIcon'
import type { Plugin } from '../../type'

import './PluginListProp.css'

const DragHandle = SortableHandle(() => <FontIcon value="menu" />)

const SortablePlugin = SortableElement(({ plugin }) =>
  <ListItem caption={plugin.name} ripple={false} leftIcon={<DragHandle />} />,
)

const SortablePluginList = SortableContainer(({ plugins }) =>
  <List>
    {plugins.map((plugin, index) =>
      <SortablePlugin key={plugin.elementName} index={index} plugin={plugin} />,
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

  render() {
    const { value } = this.props
    return (
      <SortablePluginList
        helperClass="sortablePlugin"
        plugins={value}
        onSortEnd={this.onSortEnd}
        useDragHandle={true}
      />
    )
  }
}

export default PluginListProp
