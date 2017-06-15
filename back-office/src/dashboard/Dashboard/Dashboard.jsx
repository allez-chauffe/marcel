//@flow
import React from 'react'
import Grid from '../Grid'
import { values } from 'lodash'
import type { Dashboard as DashboardT } from '../type'
import './Dashboard.css'

class Dashboard extends React.Component {
  props: { dashboard: DashboardT, deletePlugin: () => void }

  componentDidMount() {
    document.addEventListener('keydown', this.onKeyDown)
  }

  onKeyDown = ({ code }: KeyboardEvent) => {
    if (code === 'Delete' || code === 'Backspace') this.props.deletePlugin()
  }

  render() {
    const plugins = values(this.props.dashboard.plugins)
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
}

export default Dashboard
