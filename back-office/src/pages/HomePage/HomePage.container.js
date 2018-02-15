//@flow
import { compose } from 'redux'
import { connect } from 'react-redux'
import { replace } from 'redux-little-router'
import router from 'hoc-little-router'
import HomePage from './HomePage'

export default compose(
  router('HOME', { absolute: true }),
  connect(null, { goToMedias: () => replace('/medias') }),
)(HomePage)
