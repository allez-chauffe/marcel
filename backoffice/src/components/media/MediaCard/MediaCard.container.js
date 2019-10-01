import { connect } from 'react-redux'
import { compose } from 'redux'
import { withRouter } from 'react-router-dom'
import MediaCard from './MediaCard'
import { selectDashboard } from '../../../dashboard'

const mapDispatchToProps = (dispatch, ownProps) => ({
  selectDashboard: () => {
    dispatch(selectDashboard(ownProps.dashboard))
    return ownProps.history.push(`/medias/${ownProps.dashboard.id}`)
  },
})

export default compose(
  withRouter,
  connect(
    null,
    mapDispatchToProps,
  ),
)(MediaCard)
