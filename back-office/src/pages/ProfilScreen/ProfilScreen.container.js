import { connect } from 'react-redux'
import { compose } from 'redux'
import router from 'hoc-little-router'
import ProfilScreen from './ProfilScreen'
import { updateConnectedUser, updateConnectedUserProperty } from '../../auth'

const mapStateToProps = state => ({
  user: state.auth.form,
})

const mapDispatchToProps = (dispatch, ownProps) => ({
  handleSave: () => dispatch(updateConnectedUser(ownProps.user)),
  handleChange: (name, value) => dispatch(updateConnectedUserProperty(name, value)),
})

export default compose(
  router('PROFIL', { absolute: true }),
  connect(
    mapStateToProps,
    mapDispatchToProps,
  ),
)(ProfilScreen)
