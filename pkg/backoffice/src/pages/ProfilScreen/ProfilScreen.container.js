import { connect } from 'react-redux'
import { compose } from 'redux'
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
  connect(
    mapStateToProps,
    mapDispatchToProps,
  ),
)(ProfilScreen)
