import { connect } from 'react-redux'
import { updatePlugin } from '../actions'
import { pluginUpdatingSelector } from '../selectors'
import PluginCard from './PluginCard'

const mapState = state => ({
  updating: pluginUpdatingSelector(state),
})

const mapDispatch = (dispatch, { plugin: { eltName } }) => ({
  update: () => dispatch(updatePlugin(eltName)),
})

export default connect(
  mapState,
  mapDispatch,
)(PluginCard)
