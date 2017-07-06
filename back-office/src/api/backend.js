//@flow
import { keyBy } from 'lodash'

import type { Dashboard, DashboardMap } from '../dashboard/type'

const baseUrl = 'http://localhost:8090/v1/'

const request = (url: string, options?) => fetch(baseUrl + url, options)

const get = (url: string) => request(url)

const post = (url: string, body: ?mixed) =>
  request(url, {
    method: 'POST',
    headers: body ? { 'Content-Type': 'application/json' } : {},
    body: body ? JSON.stringify(body) : null,
  })

class Backend {
  getAllDashboards = (): Promise<DashboardMap> =>
    get('medias')
      .then(response => {
        if (response.status !== 200) throw response
        return response.json()
      })
      .then((dashboards: Dashboard[]) => keyBy(dashboards, 'id'))

  getDashboard = (dashboardId: string): Promise<Dashboard> =>
    get(`medias/${dashboardId}`).then(response => {
      if (response.status !== 200) throw response
      return response.json()
    })

  createDashboard = (): Promise<Dashboard> =>
    get('medias/create').then(response => {
      if (response.status !== 200) throw response
      return response.json()
    })

  saveDashboard = (dashboard: Dashboard): Promise<void> =>
    post(`medias/${dashboard.id}`, dashboard).then(response => {
      if (response.status !== 200) throw response
    })
}

export default Backend
