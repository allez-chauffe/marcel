import { connect } from 'react-redux'
import {
  login,
  changeLogin,
  changePassword,
  resetForm,
  loginSelector,
  passwordSelector,
} from '../../../auth'
import LoginForm from './LoginForm'

const mapStateToProps = state => ({
  login: loginSelector(state),
  password: passwordSelector(state),
})

const mapDispatchToProps = {
  changeLogin,
  changePassword,
  loginAction: login,
  reset: resetForm,
}

export default connect(
  mapStateToProps,
  mapDispatchToProps,
)(LoginForm)
