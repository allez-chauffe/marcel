//@flow
import { connect } from 'react-redux'
import { compose } from 'redux'
import loader from 'hoc-react-loader'
import type { Dispatch } from 'redux'
import { push } from 'redux-little-router'
import { unselectDashboard, selectedDashboardSelector } from '../../dashboard'
import { loadConfig, isConfigLoading } from '../../store/loaders'
import { LoadingIndicator } from '../../components/commons'
import AppLayout from './AppLayout'
import { logout } from '../../auth'

const mapStateToProps = state => ({
  menuIcon: selectedDashboardSelector(state) ? 'arrow_back' : null,
  loaded: !isConfigLoading(state),
  user: state.auth.user,
})

const mapDispatchToProps = (dispatch: Dispatch<*>) => ({
  goBack() {
    // WORKAROUND: Waiting for redux refactoring
    dispatch(unselectDashboard)
    dispatch(push('/medias'))
  },
  load() {
    dispatch(loadConfig())
  },
  logout() {
    dispatch(logout())
  },
})

export default compose(
  connect(mapStateToProps, mapDispatchToProps),
  loader({ print: ['loaded'], LoadingIndicator }),
)(AppLayout)
