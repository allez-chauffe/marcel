import { connect } from 'react-redux'
import { push } from 'redux-little-router'
import MediaCard from './MediaCard'
import { selectDashboard } from '../../../dashboard'

const mapDispatchToProps = (dispatch, ownProps) => ({
  selectDashboard: () => {
    dispatch(selectDashboard(ownProps.dashboard))
    dispatch(push(`/medias/${ownProps.dashboard.id}`))
  },
})

export default connect(
  null,
  mapDispatchToProps,
)(MediaCard)
