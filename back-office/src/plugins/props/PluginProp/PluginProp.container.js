import { connect } from 'react-redux'
import PluginProp from './PluginProp'
import { changeProp } from '../../../dashboard'

const mapStateToProps = state => ({
  state: state,
})

const mapDispatchToProps = (dispatch, props) => ({
  changeProp: value => dispatch(changeProp(props.plugin, props.prop, value)),
})

export default connect(
  mapStateToProps,
  mapDispatchToProps,
)(PluginProp)
