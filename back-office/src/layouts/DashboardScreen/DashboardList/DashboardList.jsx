//@flow
import React from 'react'
import type { Dashboard } from '../../../dashboard/type'
import DashboardListItem from './DashboardListItem'
import AddDashboardCard from './AddDashboardCard'

import './DashboardList.css'

export type PropsType = {
  dashboards: Dashboard[],
}

const DashboardList = (props: PropsType) => {
  const { dashboards } = props
  return (
    <div className="DashboardList">
      <AddDashboardCard />
      {dashboards.map(dashboard =>
        <DashboardListItem key={dashboard.name} dashboard={dashboard} />,
      )}
    </div>
  )
}

export default DashboardList
