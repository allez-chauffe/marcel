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
export const backendFetcher = config => {
  if (!backendFetcherInstance) {
    if (!config.urls.backend)
      toast.error("L'URL du serveur n'est pas configurée", { autoClose: false })
    const baseUrl = `${config.ssl ? 's' : ''}://${config.urls.backend}`
    backendFetcherInstance = fetcher('http' + baseUrl)
    backendFetcherInstance.ws = clientId => new WebSocket(`ws${baseUrl}/clients/${clientId}/ws`)
  }
  return backendFetcherInstance
}

let authFetcherInstance
export const authFetcher = config => {
  if (!authFetcherInstance) {
    if (!config.urls.auth)
      toast.error("L'URL du serveur d'authentification n'est pas configurée", { autoClose: false })
    const baseUrl = `http${config.ssl ? 's' : ''}://${config.urls.auth}`
    authFetcherInstance = fetcher(baseUrl)
  }
  return authFetcherInstance
}

export const localFetcher = fetcher('./')
