//@flow
import { connect } from 'react-redux'
import type { State } from '../../../store'
import {
  deletingDashboardSelector,
  confirmDashboardDeletion,
  cancelDashboardDeletion,
} from '../../../dashboard'
import DeleteDashboardDialog from './DeleteDashboardDialog'

const mapStateToProps = (state: State) => ({
  isDeletingDashboard: !!deletingDashboardSelector(state),
})

const mapDispatchToProps = {
  confirmDeletion: confirmDashboardDeletion,
  cancelDeletion: cancelDashboardDeletion,
}

export default connect(mapStateToProps, mapDispatchToProps)(DeleteDashboardDialog)
