//@flow
import React from 'react'
import ListItem from 'react-toolbox/lib/list/ListItem'
import type { Dashboard } from '../../../../dashboard/type'

export type PropsType = {
  dashboard: Dashboard,
  selectDashboard: () => void,
}

const DashboardListItem = (props: PropsType) => {
  const { dashboard, selectDashboard } = props
  return (
    <ListItem
      caption={dashboard.name}
      legend={dashboard.description}
      key={dashboard.name}
      onClick={selectDashboard}
    />
  )
}

export default DashboardListItem
