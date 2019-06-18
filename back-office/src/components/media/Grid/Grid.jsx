import React from 'react'
import { range } from 'lodash'
import ReactGridLayout from 'react-grid-layout'

import './Grid.css'
import 'react-grid-layout/css/styles.css'

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

const Grid = props => {
  const {
    size: { width, height },
    screenRatio,
    rows,
    cols,
    displayGrid,
  } = props
  const { layout, saveLayout, selectPlugin, selectedPlugin } = props
  const marginHeight = ReactGridLayout.defaultProps.margin[1]

  const containerRatio = width / height
  const gridWidth = containerRatio >= screenRatio ? screenRatio * height : width
  const gridHeight = containerRatio >= screenRatio ? height : width / screenRatio
  const rowHeight = (gridHeight - (rows + 1) * marginHeight) / rows

  return (
    <div className="Grid">
      {displayGrid ? (
        <div className="displayGrid" style={{ width: gridWidth, height: gridHeight - 10 }}>
          {range(rows).map(i => (
            <div className="row" key={i}>
              {range(cols).map(j => (
                <div className="cell" key={j} />
              ))}
            </div>
          ))}
        </div>
      ) : null}
      <ReactGridLayout
        style={{ width: gridWidth }}
        containerPadding={[5, 5]}
        width={gridWidth}
        cols={cols}
        rowHeight={rowHeight}
        verticalCompact={false}
        maxRows={rows}
        isRearrangeable={false}
        onLayoutChange={saveLayout}
      >
        {layout.map(makePluginInstance(selectPlugin, selectedPlugin))}
      </ReactGridLayout>
    </div>
  )
}

export default Grid
