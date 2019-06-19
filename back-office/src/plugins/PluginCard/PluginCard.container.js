import { connect } from 'react-redux'
import { updatePlugin } from '../actions'
import { pluginUpdating } from '../selectors'
import PluginCard from './PluginCard'

const mapState = state => ({
  updating: pluginUpdating(state),
})

const mapDispatch = (dispatch, { plugin: { eltName } }) => ({
  update: () => dispatch(updatePlugin(eltName)),
})

export default connect(
  mapState,
  mapDispatch,
)(PluginCard)
