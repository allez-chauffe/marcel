//@flow
import { connect } from 'react-redux'
import sizeMe from 'react-sizeme'
import Grid from './Grid'
import { selectPlugin, saveLayout, selectedPluginNameSelector } from '../../../dashboard'

const SizedGrid = sizeMe({ monitorWidth: true, monitorHeight: true })(Grid)

const mapStateToProps = state => ({
  selectedPlugin: selectedPluginNameSelector(state),
})

const mapDispatchToProps = {
  selectPlugin,
  saveLayout,
}

export default connect(mapStateToProps, mapDispatchToProps)(SizedGrid)
