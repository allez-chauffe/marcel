//@flow
import { connect } from 'react-redux'
import sizeMe from 'react-sizeme'
import Grid from './Grid'
import { selectedPluginNameSelector } from '../selectors'
import { selectPlugin } from '../actions'

const SizedGrid = sizeMe({ monitorWidth: true, monitorHeight: true })(Grid)

const mapStateToProps = state => ({
  selectedPlugin: selectedPluginNameSelector,
})

const mapDispatchToProps = {
  selectPlugin,
}

export default connect(mapStateToProps, mapDispatchToProps)(SizedGrid)
