import { connect } from 'react-redux'
import { addDashboard } from '../../../dashboard'
import AddMediaCard from './AddMediaCard'

const mapDispatchToProps = {
  addDashboard,
}

export default connect(
  null,
  mapDispatchToProps,
)(AddMediaCard)
