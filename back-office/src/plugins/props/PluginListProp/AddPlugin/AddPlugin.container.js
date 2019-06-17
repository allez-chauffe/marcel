import { connect } from 'react-redux'
import { pluginsSelector } from '../../../../plugins'
import AddPlugin from './AddPlugin'

const mapStateToProps = state => ({
  availablePlugins: pluginsSelector(state),
})

export default connect(mapStateToProps)(AddPlugin)
