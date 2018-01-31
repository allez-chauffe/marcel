//@flow
import React from 'react'
import Button from 'react-toolbox/lib/button/Button'
import Grid from '../Grid'
import { values } from 'lodash'

import { ActivationButton, OpenButton, DeleteDashboardButton } from '../../../common'
import type { Dashboard as DashboardT } from '../../../dashboard/type'

import './Media.css'

export type PropsType = {
  dashboard: DashboardT,
  uploadLayout: () => void,
}

const Media = (props: PropsType) => {
  const { dashboard, uploadLayout } = props
  const { name, rows, cols, screenRatio, plugins } = dashboard
  return (
    <div className="Media">
      <div className="head">
        <h2>
          {name} <br />
        </h2>
        <div className="actions">
          <Button label="Sauvegarder" icon="save" primary onClick={uploadLayout} />
          <OpenButton dashboard={dashboard} />
          <DeleteDashboardButton dashboard={dashboard} />
          <ActivationButton dashboard={dashboard} />
        </div>
      </div>
      <Grid
        screenRatio={screenRatio}
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
