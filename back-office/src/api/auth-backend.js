import store from '../store'

const baseUrl = () => store.getState().config.authURI

const request = (url, options, isJson = true) =>
  fetch(baseUrl() + url, { ...options, credentials: 'include' })
    .then(response => {
      if (~~(response.status / 100) !== 2) throw response
      return response
    })
    .then(response => (isJson ? response.json() : response))

const post = (url, body) =>
  request(url, {
    method: 'POST',
    headers: body ? { 'Content-Type': 'application/json' } : {},
    body: body ? JSON.stringify(body) : null,
  })

const put = (url, isJson) => request(url, { method: 'PUT' }, isJson)

const authBackend = {
  login: (login: ?string, password: ?string) =>
    post('login', login && password ? { login, password } : null),
  logout: () => put('logout', false),
}

export default authBackend
