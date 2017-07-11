//@flow
import { connect } from 'react-redux'
import { selectedDashboardSelector } from '../../dashboard'
import { isDashboardLoading, loadInitData } from '../../store/loaders'
import DashboardScreen from './DashboardScreen'

const mapStateToProps = state => ({
  isDashboardSelected: !!selectedDashboardSelector(state),
  isLoading: isDashboardLoading(state),
})

const mapDispatchToProps = {
  loadInitData,
}
export default connect(mapStateToProps, mapDispatchToProps)(DashboardScreen)
