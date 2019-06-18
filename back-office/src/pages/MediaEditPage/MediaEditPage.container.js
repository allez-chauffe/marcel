import { compose } from 'redux'
import { connect } from 'react-redux'
import loader from 'hoc-react-loader'
import { selectedDashboardSelector, selectDashboard } from '../../dashboard'
import { loadInitData } from '../../store/loaders'
import MediaEditPage from './MediaEditPage'

const mapStateToProps = state => ({
  media: selectedDashboardSelector(state),
})

const mapDisaptchToProps = {
  load: loadInitData,
}

// WORKAROUND: Waiting for refactoring of redux store
const workaround = connect(
  null,
  (dispatch, ownProps) => ({
    selectMedia: () => dispatch(selectDashboard(ownProps.media)),
  }),
)

export default compose(
  connect(
    mapStateToProps,
    mapDisaptchToProps,
  ),
  loader({ print: ['media'] }),
  workaround,
)(MediaEditPage)
