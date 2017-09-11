//@flow
import { connect } from 'react-redux'
import type { Dispatch } from 'redux'
import type { Dashboard } from '../../../../dashboard/type'
import { selectDashboard, requireDashboardDeletion } from '../../../../dashboard'
import type { State } from '../../../../store'
import DashboarItem from './DashboardListItem'

const mapStateToProps = (state: State) => ({
  frontendURI: state.config.frontendURI,
})

const mapDispatchToProps = (dispatch: Dispatch<*>, ownProps: { dashboard: Dashboard }) => ({
  selectDashboard: () => dispatch(selectDashboard(ownProps.dashboard)),
  deleteDashboard: () => dispatch(requireDashboardDeletion(ownProps.dashboard)),
})

export default connect(mapStateToProps, mapDispatchToProps)(DashboarItem)
