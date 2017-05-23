// @flow
import React from 'react'
import ReactGridLayout from 'react-grid-layout'

import './Grid.css'
import 'react-grid-layout/css/styles.css'

type LayoutItem = {
  x: number,
  y: number,
  w: number,
  h: number,
  id: string,
  plugin: {
    name: string,
  },
}

const Grid = ({
  size: { width, height },
  ratio = 2,
  rows = 20,
  cols = 20,
  layout = [
    { x: 0, y: 0, w: 3, h: 3, id: 'p1', plugin: { name: 'Plugin 1' } },
    { x: 1, y: 10, w: 3, h: 3, id: 'p2', plugin: { name: 'Plugin 2' } },
    { x: 5, y: 0, w: 3, h: 3, id: 'p3', plugin: { name: 'Plugin 3' } },
  ],
}: {
  size: { height: number, width: number },
  ration: number,
  ratio: number,
  rows: number,
  cols: number,
  layout: LayoutItem[],
}) => {
  const containerRatio = width / height
  const gridWidth = containerRatio >= ratio ? ratio * height : width
  const gridHeight = containerRatio >= ratio ? height : width / ratio
  const marginHeight = ReactGridLayout.defaultProps.margin[1]
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
        {layout.map(({ x, y, w, h, id, plugin }) => (
          <div key={id} data-grid={{ x, y, h, w }}>{plugin.name}</div>
        ))}
      </ReactGridLayout>
    </div>
  )
}

export default Grid
