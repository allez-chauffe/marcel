//@flow
import { connect } from 'react-redux'
import { selectedPluginSelector } from '../../../dashboard'
import SubPluginProps from './SubPluginProps'

const mapStateToProps = state => {
  return {
    plugin: selectedPluginSelector(state),
  }
}

const mapDispatchToProps = {}

export default connect(mapStateToProps, mapDispatchToProps)(SubPluginProps)
