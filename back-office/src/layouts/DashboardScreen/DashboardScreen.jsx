//@flow
import React from 'react'
import DashboardModification from './DashboardModification'
import DashboardList from './DashboardList'
import ProgressBar from 'react-toolbox/lib/progress_bar/ProgressBar'

import './DashboardScreen.css'

export type PropsType = {
  isDashboardSelected: boolean,
  isLoading: boolean,
  loadInitData: () => void,
}

class DashboardScreen extends React.Component {
  props: PropsType

  componentWillMount() {
    this.props.loadInitData()
  }

  render() {
    const { isDashboardSelected, isLoading } = this.props
    if (isLoading)
      return (
        <ProgressBar type="circular" mode="indeterminate" className="loader" />
      )
    return isDashboardSelected ? <DashboardModification /> : <DashboardList />
  }
}

export default DashboardScreen
