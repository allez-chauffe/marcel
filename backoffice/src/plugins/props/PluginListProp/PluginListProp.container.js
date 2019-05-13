import { connect } from 'react-redux'
import { pluginsSelector } from '../../selectors'
import { addSubPlugin, deletePlugin, reorderSubPlugins } from '../../../dashboard'
import PluginListProp from './PluginListProp'

const mapStateToProps = state => ({
  plugins: pluginsSelector(state),
})

const mapDispatchToProps = {
  addSubPlugin,
  deletePlugin,
  reorderSubPlugins,
}

export default connect(
  mapStateToProps,
  mapDispatchToProps,
)(PluginListProp)
