// @flow
import { connect } from 'react-redux'
import sizeMe from 'react-sizeme'
import { selectedPluginNameSelector } from '../selectors'
import { selectPlugin } from '../actions'
import Grid from './Grid'

const SizedGrid = sizeMe({ monitorWidth: true, monitorHeight: true })(Grid)

const mapStateToProps = state => ({
  selectedPlugin: selectedPluginNameSelector(state),
})

const mapDispatchToProps = {
  selectPlugin,
}

export default connect(mapStateToProps, mapDispatchToProps)(SizedGrid)
