//@flow
import React from 'react'
import type { Dashboard } from '../../../../dashboard/type'

export type PropsType = {
  dashboard: Dashboard,
  selectDashboard: () => void,
}

const DashboardListItem = (props: PropsType) => {
  const { dashboard, selectDashboard } = props
  return (
    <div key={dashboard.name} onClick={selectDashboard}>
      {dashboard.name}
    </div>
  )
}

export default DashboardListItem
