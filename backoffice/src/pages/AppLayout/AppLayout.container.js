import { connect } from 'react-redux'
import { compose } from 'redux'
import loader from 'hoc-react-loader'

import { loadConfig, isConfigLoading } from '../../store/loaders'
import { LoadingIndicator } from '../../components/commons'

import AppLayout from './AppLayout'

const mapStateToProps = state => ({
  loaded: !isConfigLoading(state),
})

const mapDispatchToProps = dispatch => ({
  load() {
    dispatch(loadConfig())
  },
})

export default compose(
  connect(
    mapStateToProps,
    mapDispatchToProps,
  ),
  loader({ print: ['loaded'], LoadingIndicator }),
)(AppLayout)
