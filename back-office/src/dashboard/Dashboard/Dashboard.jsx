//@flow
import React from 'react'
import Button from 'react-toolbox/lib/button/Button'
import Grid from '../Grid'
import { values } from 'lodash'
import type { Dashboard as DashboardT } from '../type'
import './Dashboard.css'

export type PropsType = {
  dashboard: DashboardT,
  uploadLayout: () => void,
}

const Dashboard = (props: PropsType) => {
  const { dashboard: { plugins, name, description }, uploadLayout } = props
  return (
    <div className="Dashboard">
      <div className="head">
        <h2>
          {name} <br />
          <small>{description}</small>
        </h2>
        <Button
          label="Enregistrer"
          icon="save"
          raised
          primary
          onClick={uploadLayout}
        />
      </div>
      <Grid
        ratio={2}
        rows={20}
        cols={20}
        layout={values(plugins).map(({ x, y, columns, rows, ...instance }) => ({
          layout: { x, y, h: rows, w: columns },
          plugin: instance,
        }))}
      />
    </div>
  )
}

export default Dashboard
