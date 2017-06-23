//@flow
import { connect } from 'react-redux'
import { selectedDashboardSelector } from '../selectors'
import { uploadLayout } from '../actions'
import Dashboard from './Dashboard'

const mapStateToProps = state => {
  const dashboard = selectedDashboardSelector(state)
  if (!dashboard)
    throw new Error('Illegal state ! A dashboard should be selected !')

  return { dashboard }
}

const mapDispatchToProps = {
  uploadLayout,
}

export default connect(mapStateToProps, mapDispatchToProps)(Dashboard)
