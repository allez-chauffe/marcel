//@flow
import React from 'react'
import Grid from '../Grid'
import { values } from 'lodash'
import type { Dashboard as DashboardT } from '../type'
import './Dashboard.css'

const Dashboard = (props: { dashboard: DashboardT }) => {
  const plugins = values(props.dashboard.plugins)
  return (
    <div className="Dashboard">
      <Grid
        ratio={2}
        rows={20}
        cols={20}
        layout={plugins.map(({ x, y, columns, rows, ...instance }) => ({
          layout: { x, y, h: rows, w: columns },
          plugin: instance,
        }))}
      />
    </div>
  )
}

export default Dashboard
