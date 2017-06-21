// @flow
import { connect } from 'react-redux'
import PluginProps from './PluginProps'
import type { State, Dispatch } from '../../../store'
import type { Plugin } from '../../type'
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

const mapDispatchToProps = (dispatch: Dispatch) => ({
  changeFilter: (filter: string) => dispatch(changePropsFilter(filter)),
  deletePlugin: (plugin: Plugin) => dispatch(deletePlugin(plugin)),
})

export default connect(mapStateToProps, mapDispatchToProps)(PluginProps)
