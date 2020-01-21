import { toast } from 'react-toastify'

export const fetcher = baseUrl => {
  const request = (url, options) =>
    fetch(baseUrl + url, { ...options, credentials: 'include' })
      .then(response => {
        if (~~(response.status / 100) !== 2) throw response
        return response
      })
      .then(res => res.json())

  const requestWithBody = (url, body, options = {}) =>
    request(url, {
      headers: body ? { 'Content-Type': 'application/json' } : {},
      body: body ? JSON.stringify(body) : null,
      ...options,
    })

  return {
    get: url => request(url),
    post: (url, body) => requestWithBody(url, body, { method: 'POST' }),
    put: (url, body) => requestWithBody(url, body, { method: 'PUT' }),
    del: url => request(url, { method: 'DELETE' }),
  }
}

let backendFetcherInstance
export const backendFetcher = ({ API }) => {
  if (!backendFetcherInstance) {
    if (!API) {
      toast.error("L'URL du serveur n'est pas configurée", { autoClose: false })
      return
    }

    backendFetcherInstance = fetcher(API)
    // FIXME assuming webSocket is on same domain
    const baseURI = new URL(document.baseURI)
    backendFetcherInstance.ws = clientId => new WebSocket(`${baseURI.protocol.endsWith('s:') ? 'wss' : 'ws'}://${baseURI.host}${API}clients/${clientId}/ws`)
  }
  return backendFetcherInstance
}

let authFetcherInstance
export const authFetcher = ({ API }) => {
  if (!authFetcherInstance) {
    if (!API) {
      toast.error("L'URL du serveur n'est pas configurée", { autoClose: false })
      return
    }

    authFetcherInstance = fetcher(API + 'auth/')
  }
  return authFetcherInstance
}

export const rootFetcher = fetcher('')

export const localFetcher = fetcher('./')
