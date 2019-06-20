import React from 'react'
import Button from 'react-toolbox/lib/button/Button'
import Grid from '../Grid'
import { values } from 'lodash'

import { ActivationButton, OpenButton, DeleteDashboardButton } from '../../../common'

import './Media.css'

const Media = props => {
  const { dashboard, uploadLayout } = props
  const { name, rows, cols, screenRatio, plugins, displayGrid, isWritable } = dashboard
  return (
    <div className="Media">
      <div className="head">
        <h2>
          {name} <br />
        </h2>
        <div className="actions">
          {isWritable && <Button label="Sauvegarder" icon="save" primary onClick={uploadLayout} />}
          <OpenButton dashboard={dashboard} />
          <DeleteDashboardButton dashboard={dashboard} />
          <ActivationButton dashboard={dashboard} />
        </div>
      </div>
      <Grid
        screenRatio={screenRatio}
        displayGrid={displayGrid}
        rows={rows}
        cols={cols}
        layout={values(plugins).map(({ x, y, cols, rows, ...instance }) => ({
          layout: { x, y, h: rows, w: cols },
          plugin: instance,
        }))}
      />
    </div>
  )
}

export default Media
