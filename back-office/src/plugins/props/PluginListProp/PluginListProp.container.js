//@flow
import { connect } from 'react-redux'
import { pluginsSelector } from '../../selectors'
import PluginListProp from './PluginListProp'

const mapStateToProps = (state: State) => ({
  plugins: pluginsSelector(state),
})

export default connect(mapStateToProps)(PluginListProp)
