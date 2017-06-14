//@flow
import React from 'react'
import ReactGridLayout from 'react-grid-layout'

import './Grid.css'
import 'react-grid-layout/css/styles.css'

import type { LayoutItem } from 'react-grid-layout/build/utils.js.flow'
import type { PluginInstance as PluginInstanceT } from '../type'

export type Item = {
  layout: LayoutItem,
  plugin: PluginInstanceT,
}

export type Props = {
  size: { height: number, width: number },
  ratio: number,
  rows: number,
  cols: number,
  layout: Item[],
  selectPlugin: PluginInstanceT => void,
  selectedPlugin: string,
}

const makePluginInstance = (selectPlugin, selectedPlugin) => item => {
  const { layout, plugin } = item
  const isSelected = selectedPlugin === plugin.instanceId
  return (
    <div
      key={plugin.instanceId}
      data-grid={layout}
      className={isSelected ? 'selected' : ''}
      onClick={() => selectPlugin(plugin)}
    >
      {plugin.name}
    </div>
  )
}

const Grid = (props: Props) => {
  const { size: { width, height }, ratio, rows, cols } = props
  const { layout, selectPlugin, selectedPlugin } = props
  const marginHeight: number = ReactGridLayout.defaultProps.margin[1]

  const containerRatio = width / height
  const gridWidth = containerRatio >= ratio ? ratio * height : width
  const gridHeight = containerRatio >= ratio ? height : width / ratio
  const rowHeight = (gridHeight - (rows + 1) * marginHeight) / rows

  return (
    <div className="Grid">
      <ReactGridLayout
        cols={cols}
        width={gridWidth}
        rowHeight={rowHeight}
        verticalCompact={false}
        maxRows={rows}
        isRearrangeable={false}
      >
        {layout.map(makePluginInstance(selectPlugin, selectedPlugin))}
      </ReactGridLayout>
    </div>
  )
}

export default Grid
