import { connect } from 'react-redux'
import { selectedPluginSelector, selectPluginParent } from '../../../dashboard'
import SubPluginProps from './SubPluginProps'

const mapStateToProps = state => {
  return {
    plugin: selectedPluginSelector(state),
  }
}

const mapDispatchToProps = { goBack: selectPluginParent }

export default connect(
  mapStateToProps,
  mapDispatchToProps,
)(SubPluginProps)
