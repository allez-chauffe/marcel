//@flow
import { connect } from 'react-redux'
import { compose } from 'redux'
import router from 'hoc-little-router'
import ProfilScreen from './ProfilScreen'
import { updateConnectedUser, updateConnectedUserProperty } from '../../auth'


const mapStateToProps = state => ({
  user: state.auth.user
})

const mapDispatchToProps = {
  updateUser: updateConnectedUser,
  updateUserProperty: updateConnectedUserProperty,
}
export default compose(
  router('PROFIL', { absolute: true }),  
  connect(mapStateToProps, mapDispatchToProps),
)(ProfilScreen)
