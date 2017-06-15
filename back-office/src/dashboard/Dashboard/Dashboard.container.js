//@flow
import { connect } from 'react-redux'
import { dashboardSelector } from '../selectors'
import Dashboard from './Dashboard'

const mapStateToProps = state => ({
  dashboard: dashboardSelector(state),
})

export default connect(mapStateToProps)(Dashboard)
