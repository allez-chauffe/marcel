//@flow
import { connect } from 'react-redux'
import PluginListItem from './PluginListItem'
import { addPlugin } from '../../../dashboard'

const mapDispatchToProps = {
  addPlugin,
}

export default connect(null, mapDispatchToProps)(PluginListItem)
