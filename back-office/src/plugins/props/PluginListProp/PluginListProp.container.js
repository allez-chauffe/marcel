//@flow
import { connect } from 'react-redux'
import { pluginsSelector } from '../../selectors'
import { addSubPlugin } from '../../../dashboard'
import PluginListProp from './PluginListProp'

const mapStateToProps = (state: State) => ({
  plugins: pluginsSelector(state),
})

const mapDispatchToProps = {
  addSubPlugin,
}

export default connect(mapStateToProps, mapDispatchToProps)(PluginListProp)
