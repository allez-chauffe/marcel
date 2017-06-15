//@flow
import { connect } from 'react-redux'
import PluginProp from './PluginProp'
import { changeProp } from '../../../dashboard'
import type { DashboardAction, PluginInstance } from '../../../dashboard'
import type { State } from '../../../store'
import type { Prop } from '../../../plugins'

const mapStateToProps = (state: State) => ({
  state: state,
})

const mapDispatchToProps = (
  dispatch: DashboardAction => mixed,
  props: { plugin: PluginInstance, prop: Prop },
) => ({
  changeProp: (value: mixed) =>
    dispatch(changeProp(props.plugin, props.prop, value)),
})

export default connect(mapStateToProps, mapDispatchToProps)(PluginProp)
