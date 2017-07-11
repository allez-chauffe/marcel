//@flow
import { connect } from 'react-redux'
import { selectedDashboardSelector } from '../../dashboard'
import { isLoadingInitData, loadInitData } from '../../store/loaders'
import DashboardScreen from './DashboardScreen'

const mapStateToProps = state => ({
  isDashboardSelected: !!selectedDashboardSelector(state),
  isLoading: isLoadingInitData(state),
})

const mapDispatchToProps = {
  loadInitData,
}
export default connect(mapStateToProps, mapDispatchToProps)(DashboardScreen)
