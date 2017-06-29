//@flow
import React from 'react'
import {
  SortableElement,
  SortableContainer,
  SortableHandle,
  arrayMove,
} from 'react-sortable-hoc'
import type { Plugin } from '../../type'

export type PropsType = {
  plugins: Plugin[],
  name: string,
  value: Plugin[],
  onChange: (Plugin[]) => void,
}

const SortablePlugin = SortableElement(({ plugin }) =>
  <li>
    {plugin.name}
  </li>,
)

const SortablePluginList = SortableContainer(({ plugins }) =>
  <ol>
    {plugins.map((plugin, index) =>
      <SortablePlugin key={plugin.elementName} index={index} plugin={plugin} />,
    )}
  </ol>,
)

class PluginListProp extends React.Component {
  props: PropsType

  onSortEnd = (oldIndex: number, newIndex: number) => {
    this.props.onChange(arrayMove(this.props.value, oldIndex, newIndex))
  }

  render() {
    const { value } = this.props
    return <SortablePluginList plugins={value} onSortEnd={this.onSortEnd} />
  }
}

export default PluginListProp
