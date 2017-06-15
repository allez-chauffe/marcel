//@flow
import React from 'react'
import Grid from '../Grid'
import { values } from 'lodash'
import type { Dashboard as DashboardT } from '../type'
import './Dashboard.css'

const Dashboard = ({ dashboard }: { dashboard: DashboardT }) => {
  const plugins = values(dashboard.plugins)
  return (
    <Grid
      ratio={2}
      rows={20}
      cols={20}
      layout={plugins.map(({ x, y, columns, rows, ...instance }) => ({
        layout: { x, y, h: rows, w: columns },
        plugin: instance,
      }))}
    />
  )
}

export default Dashboard
