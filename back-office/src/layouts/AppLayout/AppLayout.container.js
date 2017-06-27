//@flow
import { connect } from 'react-redux'
import { unselectDashboard, selectedDashboardSelector } from '../../dashboard'
import AppLayout from './AppLayout'

const mapStateToProps = state => ({
  isDashboardSelected: !!selectedDashboardSelector(state),
})

const mapDispatchToProps = {
  unselectDashboard,
}

export default connect(mapStateToProps, mapDispatchToProps)(AppLayout)
