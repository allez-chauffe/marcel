//@flow
import { values, mapValues } from 'lodash'
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
    get('medias/create').then(response => {
      if (response.status !== 200) throw response
      return response.json()
    }),

  saveDashboard: (dashboard: Dashboard) => {
    const plugins = values(dashboard.plugins).map(plugin => {
      const { x, y, cols, rows, props, eltName, name, instanceId } = plugin
      const propsForBack = mapValues(props, 'value')
      return {
        name,
        instanceId,
        frontend: { x, y, cols, rows, eltName, props: propsForBack },
      }
    })
    const data = { ...dashboard, plugins }
    return post(`medias`, data).then(response => {
      if (response.status !== 200) throw response
    })
  },

  getAvailablePlugins: () => new Promise(resolve => resolve(availablePlugins)),
}

export default backend
