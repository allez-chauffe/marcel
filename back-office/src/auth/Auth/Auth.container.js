//@flow
import { connect } from 'react-redux'
import { isLoggedInSelector, isLoading } from '../selectors'
import { refreshLogin } from '../actions'
import type { State } from '../../store'

import Auth from './Auth'

const mapStateToProps = (state: State) => ({
  isLoggedIn: isLoggedInSelector(state),
  isLoading: isLoading(state),
})

const mapDispatchToProps = {
  login: refreshLogin,
}

export default connect(mapStateToProps, mapDispatchToProps)(Auth)
