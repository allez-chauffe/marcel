import { connect } from 'react-redux'
import { updatePlugin, deletePlugin } from '../actions'
import { pluginUpdating } from '../selectors'
import PluginCard from './PluginCard'

const mapState = (state, { plugin: { eltName } }) => ({
  updating: pluginUpdating(eltName, state),
})

const mapDispatch = (dispatch, { plugin: { eltName } }) => ({
  update: () => dispatch(updatePlugin(eltName)),
  deletePlugin: () => dispatch(deletePlugin(eltName)),
})

export default connect(
  mapState,
  mapDispatch,
)(PluginCard)
