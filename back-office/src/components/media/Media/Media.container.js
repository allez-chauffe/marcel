import { connect } from 'react-redux'
import { uploadLayout, selectedDashboardSelector } from '../../../dashboard'
import Dashboard from './Media'

const mapStateToProps = state => {
  const dashboard = selectedDashboardSelector(state)
  if (!dashboard) throw new Error('Illegal state ! A dashboard should be selected !')

  return { dashboard }
}

const mapDispatchToProps = {
  uploadLayout,
}

export default connect(
  mapStateToProps,
  mapDispatchToProps,
)(Dashboard)
