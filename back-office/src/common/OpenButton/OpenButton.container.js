//@flow
import { connect } from 'react-redux'
import OpenButton from './OpenButton'

const mapStateToProps = state => ({
  frontendURI: state.config.frontendURI,
})

export default connect(mapStateToProps)(OpenButton)
