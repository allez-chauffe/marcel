//@flow
import { connect } from 'react-redux'
import UserLine from './UserLine'
import { editUser, deleteUser } from '../../../user'


const mapDispatchToProps = {
  editUser: editUser,
  deleteUser: deleteUser,
}

export default connect(null,mapDispatchToProps)(UserLine)
