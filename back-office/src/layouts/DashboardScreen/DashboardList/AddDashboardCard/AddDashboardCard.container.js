//@flow
import { connect } from 'react-redux'
import ripple from 'react-toolbox/lib/ripple/Ripple'
import { addDashboard } from '../../../../dashboard'
import AddDashboardCard from './AddDashboardCard'

const RippledCard = ripple({ spread: 3 })(AddDashboardCard)

const mapDispatchToProps = {
  addDashboard,
}

export default connect(null, mapDispatchToProps)(RippledCard)
