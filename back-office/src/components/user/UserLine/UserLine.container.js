//@flow
import { connect } from 'react-redux'
import UserLine from './UserLine'
import { editUser, deleteUser } from '../../../user'

const mapDispatchToProps = (dispatch, ownProps) => ({
  handleEdit: () => dispatch(editUser(ownProps.user)),
  handleDelete: () => dispatch(deleteUser(ownProps.user.id))
})

export default connect(null, mapDispatchToProps)(UserLine)
