import { connect } from 'react-redux'
import DeleteButton from './DeleteDashboardButton'
import { requireDashboardDeletion } from '../../dashboard'

const mapDispatchToProps = (dispatch, ownProps) => ({
  delete: () => dispatch(requireDashboardDeletion(ownProps.dashboard)),
})

export default connect(
  null,
  mapDispatchToProps,
)(DeleteButton)
