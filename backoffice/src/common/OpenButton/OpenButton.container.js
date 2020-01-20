import { connect } from 'react-redux'
import OpenButton from './OpenButton'

const mapStateToProps = state => ({
  Frontend: state.config.Frontend,
})

export default connect(mapStateToProps)(OpenButton)
