import { connect } from 'react-redux'
import PluginProps from './PluginProps'
import { deletePlugin } from '../../../dashboard'
import {
  propsFilterSelector,
  changePropsFilter,
  selectedPluginPropsFilteredSelector,
} from '../../../store/filters'

const mapStateToProps = state => ({
  filter: propsFilterSelector(state),
  plugin: selectedPluginPropsFilteredSelector(state),
})

const mapDispatchToProps = dispatch => ({
  changeFilter: filter => dispatch(changePropsFilter(filter)),
  deletePlugin: plugin => dispatch(deletePlugin(plugin)),
})

export default connect(
  mapStateToProps,
  mapDispatchToProps,
)(PluginProps)
