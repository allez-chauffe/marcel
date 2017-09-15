//@flow
import { connect } from 'react-redux'
import type { Dispatch } from 'redux'
import DeleteButton from './DeleteDashboardButton'
import { requireDashboardDeletion } from '../../dashboard'
import type { DashboardAction } from '../../dashboard'

const mapDispatchToProps = (dispatch: Dispatch<DashboardAction>, ownProps) => ({
  delete: () => dispatch(requireDashboardDeletion(ownProps.dashboard)),
})

export default connect(null, mapDispatchToProps)(DeleteButton)
