//@flow
import React from 'react'
import List from 'react-toolbox/lib/list/List'
import type { Dashboard } from '../../../dashboard/type'
import DashboardListItem from './DashboardListItem'

export type PropsType = {
  dashboards: Dashboard[],
}

const DashboardList = (props: PropsType) => {
  const { dashboards } = props
  return (
    <List selectable>
      {dashboards.map(dashboard =>
        <DashboardListItem key={dashboard.name} dashboard={dashboard} />,
      )}
    </List>
  )
}

export default DashboardList
