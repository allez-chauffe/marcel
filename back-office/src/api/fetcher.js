//@flow
import store from '../store'
import authBackend from './auth-backend'
import { disconnected } from '../auth'

const fetcher = (baseUrl: () => string) => {
  const request = (url: string, options?) => {
    const request = () =>
      fetch(baseUrl() + url, { ...options, credentials: 'include' }).then(response => {
        if (~~(response.status / 100) !== 2) throw response
        return response
      })

    return request().catch(response => {
      if (response.status !== 403) throw response
      return authBackend
        .login()
        .then(request)
        .catch(response => {
          if (response.status !== 403) throw response
          store.dispatch(disconnected())
        })
    })
  }

  const requestWithBody = (url: string, body: ?mixed, options = {}) =>
    request(url, {
      headers: body ? { 'Content-Type': 'application/json' } : {},
      body: body ? JSON.stringify(body) : null,
      ...options,
    })

  return {
    get: (url: string) => request(url),
    post: (url: string, body: ?mixed) => requestWithBody(url, body, { method: 'POST' }),
    put: (url: string, body: ?mixed) => requestWithBody(url, body, { method: 'PUT' }),
    del: (url: string) => request(url, { method: 'DELETE' }),
  }
}

export default fetcher
