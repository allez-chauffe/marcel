import { connect } from 'react-redux'
import { compose } from 'redux'
import { withRouter } from 'react-router-dom'

import { unselectDashboard, selectedDashboardSelector } from '../../../dashboard'
import { logout } from '../../../auth'

import Header from './Header'

const mapStateToProps = state => ({
  menuIcon: selectedDashboardSelector(state) ? 'arrow_back' : null,
  user: state.auth.user,
})

const mapDispatchToProps = (dispatch, {history}) => ({
  goBack() {
    // WORKAROUND: Waiting for redux refactoring
    dispatch(unselectDashboard)
    return history.push('/medias')
  },
  logout() {
    dispatch(logout())
  },
})

export default compose(
  withRouter,
  connect(
    mapStateToProps,
    mapDispatchToProps,
  ),
)(Header)
