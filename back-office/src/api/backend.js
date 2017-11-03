//@flow
import store from '../store'
import { values, mapValues, pick, keyBy } from 'lodash'
import type { Dashboard } from '../dashboard/type'
import type { Client } from '../clients'

const baseUrl = () => store.getState().config.backendURI

const request = (url: string, options?) =>
  fetch(baseUrl() + url, options).then(response => {
    if (~~(response.status / 100) !== 2) throw response
    return response
  })

const requestWithBody = (url: string, body: ?mixed, options = {}) =>
  request(url, {
    headers: body ? { 'Content-Type': 'application/json' } : {},
    body: body ? JSON.stringify(body) : null,
    ...options,
  })

const get = (url: string) => request(url)

const post = (url: string, body: ?mixed) => requestWithBody(url, body, { method: 'POST' })

const put = (url: string, body: ?mixed) => requestWithBody(url, body, { method: 'PUT' })

const del = (url: string) => request(url, { method: 'DELETE' })

const backend = {
  getAllDashboards: () => get('medias/').then(res => res.json()),

  getDashboard: (dashboardId: string) => get(`medias/${dashboardId}/`).then(res => res.json()),

  createDashboard: () => post('medias/').then(res => res.json()),

  saveDashboard: (dashboard: Dashboard) => {
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

  updateClient: (client: Client) => put('clients/', client),

  activateDashboard: (dashboardId: string) => get(`medias/${dashboardId}/activate`),

  deactivateDashboard: (dashboardId: string) => get(`medias/${dashboardId}/deactivate`),

  deleteDashboard: (dashboardId: string) => del(`medias/${dashboardId}/`),
}

export default backend
