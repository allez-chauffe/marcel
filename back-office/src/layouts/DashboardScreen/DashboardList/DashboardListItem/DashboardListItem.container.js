//@flow
import { connect } from 'react-redux'
import type { Dispatch } from 'redux'
import type { Dashboard } from '../../../../dashboard/type'
import {
  selectDashboard,
  requireDashboardDeletion,
} from '../../../../dashboard'
import DashboarItem from './DashboardListItem'

const mapDispatchToProps = (
  dispatch: Dispatch<*>,
  ownProps: { dashboard: Dashboard },
) => ({
  selectDashboard: () => dispatch(selectDashboard(ownProps.dashboard)),
  deleteDashboard: () => dispatch(requireDashboardDeletion(ownProps.dashboard)),
})

export default connect(null, mapDispatchToProps)(DashboarItem)
