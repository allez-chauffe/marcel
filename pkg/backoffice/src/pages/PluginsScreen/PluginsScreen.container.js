import { compose } from 'redux'
import { connect } from 'react-redux'
import loader from 'hoc-react-loader'
import { LoadingIndicator } from '../../components/commons'
import PluginsScreen from './PluginsScreen'
import { loadPlugins, isPluginsLoading } from '../../store/loaders'
import { pluginsSelector } from '../../plugins'

const mapState = state => ({
  plugins: pluginsSelector(state),
  loaded: !isPluginsLoading(state),
})

const mapDispatch = {
  load: loadPlugins,
}

export default compose(
  connect(
    mapState,
    mapDispatch,
  ),
  loader({ LoadingIndicator }),
)(PluginsScreen)
