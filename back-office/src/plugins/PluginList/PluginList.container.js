//@flow
import { connect } from 'react-redux'
import PluginList from './PluginList'
import { addPlugin } from '../../dashboard'
import type { State } from '../../store'
import {
  pluginFilterSelector,
  changePluginFilter,
  filteredPluginsSeletor,
} from '../../store/filters'

const mapStateToProps = (state: State) => ({
  filter: pluginFilterSelector(state),
  plugins: filteredPluginsSeletor(state),
})

const mapDisaptchToProps = {
  changeFilter: changePluginFilter,
  addPlugin,
}

export default connect(mapStateToProps, mapDisaptchToProps)(PluginList)
