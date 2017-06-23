//@flow
import { connect } from 'react-redux'
import { values } from 'lodash'
import { dashboardsSelector } from '../../../dashboard'
import type { State } from '../../../store'

import DashboardList from './DashboardList'

const mapStateToProps = (state: State) => ({
  dashboards: values(dashboardsSelector(state)),
})

export default connect(mapStateToProps)(DashboardList)
