//@flow
import { connect } from 'react-redux'
import sizeMe from 'react-sizeme'
import Grid from './Grid'
import { selectedPluginNameSelector, displayGridSelector } from '../selectors'
import { selectPlugin, saveLayout } from '../actions'

const SizedGrid = sizeMe({ monitorWidth: true, monitorHeight: true })(Grid)

const mapStateToProps = state => ({
  selectedPlugin: selectedPluginNameSelector(state),
  displayGrid: displayGridSelector(state),
})

const mapDispatchToProps = {
  selectPlugin,
  saveLayout,
}

export default connect(mapStateToProps, mapDispatchToProps)(SizedGrid)
