// @flow
import { connect } from 'react-redux'
import PluginProps from './PluginProps'
import type { State } from '../../../store'
import { propsFilterSelector, changePropsFilter } from '../../../store/filters'

const mapStateToProps = (state: State) => ({
  filter: propsFilterSelector(state),
})

const mapDispatchToProps = {
  changeFilter: changePropsFilter,
}

export default connect(mapStateToProps, mapDispatchToProps)(PluginProps)
