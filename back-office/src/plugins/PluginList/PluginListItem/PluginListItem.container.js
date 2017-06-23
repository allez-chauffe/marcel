//@flow
import { connect } from 'react-redux'
import type { Dispatch } from 'redux'
import PluginListItem from './PluginListItem'
import { addPlugin } from '../../../dashboard'
import type { Plugin } from '../../../plugins'

const mapDispatchToProps = (dispatch: Dispatch<*>) => ({
  addPlugin: (plugin: Plugin) => dispatch(addPlugin(plugin)),
})

export default connect(null, mapDispatchToProps)(PluginListItem)
