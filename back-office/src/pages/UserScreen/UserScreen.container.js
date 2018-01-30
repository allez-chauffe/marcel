//@flow
import { connect } from 'react-redux'
import { compose } from 'redux'
import { isUsersLoading, loadUsers } from '../../store/loaders'
import router from 'hoc-little-router'
import UserScreen from './UserScreen'
import loader from 'hoc-react-loader'
import { LoadingIndicator } from '../../components/commons'
import values from 'lodash/values'
import { addUser, updateUser, updateCurrentUserProperty, resetCurrentUser } from '../../user'

const mapStateToProps = state => ({
  userEdited: state.users.currentUser,
  users: values(state.users.users),
  loaded: !isUsersLoading(state),
})

const mapDispatchToProps = {
  load: loadUsers,
  addUser: addUser,
  updateUser: updateUser,
  updateCurrentUserProperty: updateCurrentUserProperty,
  resetCurrentUser: resetCurrentUser,
}
export default compose(
  router('USERS', { absolute: true }),  
  connect(mapStateToProps, mapDispatchToProps),
  loader({ print:['loaded'], LoadingIndicator }),
)(UserScreen)
