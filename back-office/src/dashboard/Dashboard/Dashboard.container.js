//@flow
import { connect } from 'react-redux'
import { selectedDashboardSelector } from '../selectors'
import { uploadLayout } from '../actions'
import Dashboard from './Dashboard'

const mapStateToProps = state => ({
  dashboard: selectedDashboardSelector(state),
})

const mapDispatchToProps = {
  uploadLayout,
}

export default connect(mapStateToProps, mapDispatchToProps)(Dashboard)
