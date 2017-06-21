//@flow
import React from 'react'
import type { Dashboard } from '../../../dashboard/type'

export type PropsType = {
  dashboards: Dashboard[],
  selectDashboard: Dashboard => void,
}

const DashboardList = (props: PropsType) => {
  const { dashboards, selectDashboard } = props
  return (
    <div>
      {dashboards.map(dashboard =>
        <div key={dashboard.name} onClick={() => selectDashboard(dashboard)}>
          {dashboard.name}
        </div>,
      )}
    </div>
  )
}

export default DashboardList
