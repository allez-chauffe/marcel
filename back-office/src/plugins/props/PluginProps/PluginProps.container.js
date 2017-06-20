// @flow
import { connect } from 'react-redux'
import PluginProps from './PluginProps'
import type { State } from '../../../store'
import { deletePlugin } from '../../../dashboard'
import {
  propsFilterSelector,
  changePropsFilter,
  selectedPluginPropsFilteredSelector,
} from '../../../store/filters'

const mapStateToProps = (state: State) => ({
  filter: propsFilterSelector(state),
  plugin: selectedPluginPropsFilteredSelector(state),
})

const mapDispatchToProps = {
  changeFilter: changePropsFilter,
  deletePlugin,
}

export default connect(mapStateToProps, mapDispatchToProps)(PluginProps)
