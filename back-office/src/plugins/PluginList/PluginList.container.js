//@flow
import { connect } from 'react-redux'
import PluginList from './PluginList'
import { addPlugin } from '../../dashboard'
import {
  pluginFilterSelector,
  changePluginFilter,
  filteredPluginsSeletor,
} from '../../store/filters'

const mapStateToProps = state => ({
  filter: pluginFilterSelector(state),
  plugins: filteredPluginsSeletor(state),
})

const mapDisaptchToProps = {
  changeFilter: changePluginFilter,
  addPlugin,
}

export default connect(mapStateToProps, mapDisaptchToProps)(PluginList)
