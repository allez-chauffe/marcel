//@flow
import { connect } from 'react-redux'
import { addDashboard } from '../../../../dashboard'
import AddDashboardCard from './AddDashboardCard'

const mapDispatchToProps = {
  addDashboard,
}

export default connect(null, mapDispatchToProps)(AddDashboardCard)
