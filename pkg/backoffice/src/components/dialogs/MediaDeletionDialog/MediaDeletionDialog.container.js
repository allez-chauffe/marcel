import { connect } from 'react-redux'
import {
  deletingDashboardSelector,
  confirmDashboardDeletion,
  cancelDashboardDeletion,
} from '../../../dashboard'
import DeleteDashboardDialog from './MediaDeletionDialog'

const mapStateToProps = state => ({
  isDeletingDashboard: !!deletingDashboardSelector(state),
})

const mapDispatchToProps = {
  confirmDeletion: confirmDashboardDeletion,
  cancelDeletion: cancelDashboardDeletion,
}

export default connect(
  mapStateToProps,
  mapDispatchToProps,
)(DeleteDashboardDialog)
