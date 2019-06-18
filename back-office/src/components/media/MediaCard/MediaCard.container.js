import { connect } from 'react-redux'
import { navigate } from '@reach/router'
import MediaCard from './MediaCard'
import { selectDashboard } from '../../../dashboard'

const mapDispatchToProps = (dispatch, ownProps) => ({
  selectDashboard: () => {
    dispatch(selectDashboard(ownProps.dashboard))
    return navigate(`/medias/${ownProps.dashboard.id}`)
  },
})

export default connect(
  null,
  mapDispatchToProps,
)(MediaCard)
