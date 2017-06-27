//@flow
import { connect } from 'react-redux'
import { isLoggedInSelector } from '../selectors'
import { login, logout } from '../actions'
import type { State } from '../../store'

import Auth from './Auth'

const mapStateToProps = (state: State) => ({
  isLoggedIn: isLoggedInSelector(state),
})

const mapDispatchToProps = {
  login,
}

export default connect(mapStateToProps, mapDispatchToProps)(Auth)
