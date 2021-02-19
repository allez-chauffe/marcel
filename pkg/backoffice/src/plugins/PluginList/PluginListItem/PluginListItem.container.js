import { connect } from 'react-redux'
import PluginListItem from './PluginListItem'
import { addPlugin } from '../../../dashboard'

const mapDispatchToProps = dispatch => ({
  addPlugin: plugin => dispatch(addPlugin(plugin)),
})

export default connect(
  null,
  mapDispatchToProps,
)(PluginListItem)
