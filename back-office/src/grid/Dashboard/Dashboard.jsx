// @flow
import React from 'react'
import Grid from '../Grid'
import './Dashboard.css'

const Dashboard = () => (
  <Grid
    ratio={2}
    rows={20}
    cols={20}
    layout={[
      {
        layout: { x: 0, y: 0, w: 3, h: 3 },
        id: 'p1',
        plugin: { name: 'Plugin 1' },
      },
      {
        layout: { x: 1, y: 10, w: 3, h: 3 },
        id: 'p2',
        plugin: { name: 'Plugin 2' },
      },
      {
        layout: { x: 5, y: 0, w: 3, h: 3 },
        id: 'p3',
        plugin: { name: 'Plugin 3' },
      },
    ]}
  />
)

export default Dashboard
