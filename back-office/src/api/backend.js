import { values, mapValues, pick, keyBy } from 'lodash'
import store from '../store'
import fetcher from './fetcher'

const { get, post, put, del } = fetcher(() => store.getState().config.backendURI)

const backend = {
  getAllDashboards: () => get('medias/').then(res => res.json()),

  getDashboard: dashboardId => get(`medias/${dashboardId}/`).then(res => res.json()),

  createDashboard: () => post('medias/').then(res => res.json()),

  saveDashboard: dashboard => {
    const plugins = values(dashboard.plugins).map(plugin => {
      const { x, y, cols, rows, props, eltName, instanceId } = plugin
      const propsForBack = mapValues(props, 'value')
      return {
        instanceId,
        eltName,
        frontend: { x, y, cols, rows, props: propsForBack },
      }
    })
    const data = { ...dashboard, plugins }
    return put(`medias/`, data)
  },

  getAvailablePlugins: () =>
    get('plugins/')
      .then(res => res.json())
      .then(plugins =>
        plugins.map(plugin => ({
          ...pick(plugin, 'name', 'description', 'eltName'),
          props: keyBy(plugin.frontend.props, 'name'),
        })),
      ),

  getClients: () => get('clients/').then(res => res.json()),

  updateClient: client => put('clients/', client),

  activateDashboard: dashboardId => get(`medias/${dashboardId}/activate`),

  deactivateDashboard: dashboardId => get(`medias/${dashboardId}/deactivate`),

  deleteDashboard: dashboardId => del(`medias/${dashboardId}/`),
}

export default backend
