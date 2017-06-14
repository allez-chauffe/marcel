// @flow
import React from 'react'
import ReactGridLayout from 'react-grid-layout'

import './Grid.css'
import 'react-grid-layout/css/styles.css'

import type { LayoutItem } from 'react-grid-layout/build/utils.js.flow'

export type Item = {
  layout: LayoutItem,
  id: string,
  plugin: {
    name: string,
    elementName: string,
  },
}

export type Props = {
  size: { height: number, width: number },
  ration: number,
  ratio: number,
  rows: number,
  cols: number,
  layout: Item[],
  selectPlugin: string => void,
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
        {layout.map(({ layout, plugin, id }) => (
          <div
            key={id}
            data-grid={layout}
            className={selectedPlugin === plugin.elementName ? 'selected' : ''}
            onClick={() => selectPlugin(plugin.elementName)}
          >
            {plugin.name}
          </div>
        ))}
      </ReactGridLayout>
    </div>
  )
}

export default Grid
