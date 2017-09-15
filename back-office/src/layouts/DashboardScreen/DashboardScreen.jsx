//@flow
import React from 'react'
import DashboardModification from './DashboardModification'
import DashboardList from './DashboardList'
import ProgressBar from 'react-toolbox/lib/progress_bar/ProgressBar'
import DeleteDashboardDialog from './DeleteDashboardDialog'

import './DashboardScreen.css'

class DashboardScreen extends React.Component {
  props: {
    isDashboardSelected: boolean,
    isLoading: boolean,
    loadInitData: () => void,
  }

  componentWillMount() {
    this.props.loadInitData()
  }

  render() {
    const { isDashboardSelected, isLoading } = this.props
    if (isLoading) return <ProgressBar type="circular" mode="indeterminate" className="loader" />
    return (
      <div className="DashboardScreen">
        {isDashboardSelected ? <DashboardModification /> : <DashboardList />}
        <DeleteDashboardDialog />
      </div>
    )
  }
}

export default DashboardScreen
