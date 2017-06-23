//@flow
import { connect } from 'react-redux'
import { selectedDashboardSelector } from '../../dashboard'
import DashboardScreen from './DashboardScreen'

const mapStateToProps = state => ({
  isDashboardSelected: !!selectedDashboardSelector(state),
})

export default connect(mapStateToProps)(DashboardScreen)
