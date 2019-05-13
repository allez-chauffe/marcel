import { compose } from 'redux'
import loader from 'hoc-react-loader'
import { connect } from 'react-redux'
import { isLoggedInSelector, isLoading, refreshLogin } from '../../../auth'

import { LoadingIndicator } from '../../commons'
import Auth from './Auth'

const mapStateToProps = state => ({
  isLoggedIn: isLoggedInSelector(state),
  loaded: !isLoading(state),
})

const mapDispatchToProps = {
  load: refreshLogin,
}

export default compose(
  connect(
    mapStateToProps,
    mapDispatchToProps,
  ),
  loader({ print: ['loaded'], LoadingIndicator }),
)(Auth)
