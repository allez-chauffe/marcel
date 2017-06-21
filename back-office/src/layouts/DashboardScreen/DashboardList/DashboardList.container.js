//@flow
import { connect } from 'react-redux'
import { values } from 'lodash'
import { dashboardsSelector, selectDashboard } from '../../../dashboard'
import type { State } from '../../../store'

import DashboardList from './DashboardList'

const mapStateToProps = (state: State) => ({
  dashboards: values(dashboardsSelector(state)),
})

const mapDispatchToProps = {
  selectDashboard,
}

export default connect(mapStateToProps, mapDispatchToProps)(DashboardList)
