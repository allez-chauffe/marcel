// @flow
import React from 'react'
import ReactGridLayout from 'react-grid-layout'

import './Grid.css'
import 'react-grid-layout/css/styles.css'

import type { LayoutItem } from 'react-grid-layout/build/utils.js.flow'
import type { PluginInstance } from '../type'

export type Item = {
  layout: LayoutItem,
  id: string,
  plugin: PluginInstance,
}

export type Props = {
  size: { height: number, width: number },
  ration: number,
  ratio: number,
  rows: number,
  cols: number,
  layout: Item[],
  selectPlugin: PluginInstance => void,
  selectedPlugin: string,
}

const Grid = (props: Props) => {
  const {
    size,
    ratio,
    rows,
    cols,
    layout,
    selectPlugin,
    selectedPlugin,
  } = props
  const { width, height } = size
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
        {layout.map(({ layout, plugin }) => (
          <div
            key={plugin.instanceId}
            data-grid={layout}
            className={selectedPlugin === plugin.instanceId ? 'selected' : ''}
            onClick={() => selectPlugin(plugin)}
          >
            {plugin.name}
          </div>
        ))}
      </ReactGridLayout>
    </div>
  )
}

export default Grid
