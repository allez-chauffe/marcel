//@flow
import { values, mapValues, pick, keyBy } from 'lodash'
import type { Dashboard } from '../dashboard/type'
import availablePlugins from '../mocked-data/plugins'

const baseUrl = 'http://localhost:8090/api/v1/'

const request = (url: string, options?) => fetch(baseUrl + url, options)

const get = (url: string) => request(url)

const post = (url: string, body: ?mixed) =>
  request(url, {
    method: 'POST',
    headers: body ? { 'Content-Type': 'application/json' } : {},
    body: body ? JSON.stringify(body) : null,
  })

const put = (url: string, body: ?mixed) =>
  request(url, {
    method: 'PUT',
    headers: body ? { 'Content-Type': 'application/json' } : {},
    body: body ? JSON.stringify(body) : null,
  })

const backend = {
  getAllDashboards: () =>
    get('medias').then(response => {
      if (response.status !== 200) throw response
      return response.json()
    }),

  getDashboard: (dashboardId: string) =>
    get(`medias/${dashboardId}`).then(response => {
      if (response.status !== 200) throw response
      return response.json()
    }),

  createDashboard: () =>
    post('medias').then(response => {
      if (response.status !== 200) throw response
      return response.json()
    }),

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
    return put(`medias`, data).then(response => {
      if (response.status !== 200) throw response
    })
  },

  getAvailablePlugins: () =>
    get('plugins')
      .then(response => {
        if (response.status !== 200) throw response
        return response.json()
      })
      .then(plugins =>
        plugins.map(plugin => ({
          ...pick(plugin, 'name', 'description', 'eltName'),
          props: keyBy(plugin.frontend.props, 'name'),
        })),
      ),
}

export default backend
