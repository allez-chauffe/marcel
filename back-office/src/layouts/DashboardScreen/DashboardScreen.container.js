//@flow
import { connect } from 'react-redux'
import {
  selectedDashboardSelector,
  isLoadingSelector,
  loadDashboardList,
} from '../../dashboard'
import DashboardScreen from './DashboardScreen'

const mapStateToProps = state => ({
  isDashboardSelected: !!selectedDashboardSelector(state),
  isLoading: isLoadingSelector(state),
})

const mapDispatchToProps = {
  loadDashboardList,
}
export default connect(mapStateToProps, mapDispatchToProps)(DashboardScreen)
