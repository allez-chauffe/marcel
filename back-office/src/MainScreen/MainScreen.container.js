// @flow
import { connect } from 'react-redux'
import MainScreen from './MainScreen'

import type { State } from '../store/types'

const mapStateToProps = (state: State) => ({
  availablePlugins: state.plugins,
})

const mapDispatchToProps = {}

export default connect(mapStateToProps, mapDispatchToProps)(MainScreen)
