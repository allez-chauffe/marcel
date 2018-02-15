//@flow
import { connect } from 'react-redux'
import { push } from 'redux-little-router'
import type { Dispatch } from 'redux'
import type { Dashboard } from '../../../dashboard/type'
import MediaCard from './MediaCard'
import { selectDashboard } from '../../../dashboard'

const mapDispatchToProps = (dispatch: Dispatch<*>, ownProps: { dashboard: Dashboard }) => ({
  selectDashboard: () => {
    dispatch(selectDashboard(ownProps.dashboard))
    dispatch(push(`/medias/${ownProps.dashboard.id}`))
  },
})

export default connect(null, mapDispatchToProps)(MediaCard)
