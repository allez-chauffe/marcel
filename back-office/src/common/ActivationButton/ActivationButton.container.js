//@flow
import { connect } from 'react-redux'
import ActivationButton from './ActivationButton'
import { activateDashboard, deactivateDashboard } from '../../dashboard/actions'
import type { DashboardAction } from '../../dashboard/type'

const mapStateToProps = (state: State, ownProps) => ({
  isActive: ownProps.dashboard.isactive,
})

const mapDispatchToProps = (dispatch: Dispatch<DashboardAction>, ownProps) => ({
  activate: () => dispatch(activateDashboard(ownProps.dashboard.id)),
  deactivate: () => dispatch(deactivateDashboard(ownProps.dashboard.id)),
})

export default connect(mapStateToProps, mapDispatchToProps)(ActivationButton)
