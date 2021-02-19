import { connect } from 'react-redux'
import ActivationButton from './ActivationButton'
import { activateDashboard, deactivateDashboard } from '../../dashboard/actions'

const mapStateToProps = (state, ownProps) => ({
  isActive: ownProps.dashboard.isactive,
  isWritable: ownProps.dashboard.isWritable,
})

const mapDispatchToProps = (dispatch, ownProps) => ({
  activate: () => dispatch(activateDashboard(ownProps.dashboard.id)),
  deactivate: () => dispatch(deactivateDashboard(ownProps.dashboard.id)),
})

export default connect(
  mapStateToProps,
  mapDispatchToProps,
)(ActivationButton)
