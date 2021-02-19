import { connect } from 'react-redux'
import AddPlugin from './AddPlugin'
import { addPlugin } from '../actions'
import { addingPlugin } from '../selectors'

const mapState = state => ({
  adding: addingPlugin(state),
})

const mapDispatch = {
  add: addPlugin,
}

export default connect(
  mapState,
  mapDispatch,
)(AddPlugin)
