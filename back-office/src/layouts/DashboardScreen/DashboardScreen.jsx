//@flow
import React from 'react'
import DashboardModification from './DashboardModification'
import DashboardList from './DashboardList'

export type PropsType = {
  isDashboardSelected: boolean,
}

const DashboardScreen = (props: PropsType) => {
  const { isDashboardSelected } = props
  return isDashboardSelected ? <DashboardModification /> : <DashboardList />
}

export default DashboardScreen
