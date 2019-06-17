import { connect } from 'react-redux'
import { compose } from 'redux'
import router from 'hoc-little-router'
import loader from 'hoc-react-loader'
import values from 'lodash/values'
import { LoadingIndicator } from '../../components/commons'
import { dashboardsSelector } from '../../dashboard'
import { isLoadingInitData, loadInitData } from '../../store/loaders'
import MediaListPage from './MediaListPage'

const mapStateToProps = state => ({
  loaded: !isLoadingInitData(state),
  medias: values(dashboardsSelector(state)),
})

const mapDispatchToProps = {
  load: loadInitData,
}

export default compose(
  router('MEDIAS', { absolute: true }),
  connect(
    mapStateToProps,
    mapDispatchToProps,
  ),
  loader({ LoadingIndicator }),
)(MediaListPage)
