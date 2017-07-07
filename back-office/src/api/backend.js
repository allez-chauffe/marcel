//@flow
import type { Dashboard } from '../dashboard/type'

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
    get('medias/create')
      .then(response => {
        if (response.status !== 200) throw response
        return response.json()
      })
      .then(dashboard => ({
        ...dashboard,
        name: 'Dashboard',
        rows: 10,
        cols: 10,
        plugiins: [],
      })),

  saveDashboard: (dashboard: Dashboard) =>
    post(`medias/${dashboard.id}`, dashboard).then(response => {
      if (response.status !== 200) throw response
    }),
}

export default backend
