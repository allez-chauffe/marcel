//@flow
import store from '../store'

const baseUrl = () => store.getState().config.authURI

const request = (url, options) =>
  fetch(baseUrl() + url, { ...options, credentials: 'include' })
    .then(response => {
      if (~~(response.status / 100) !== 2) throw response
      return response
    })
    .then(response => response.json())

const post = (url, body) =>
  request(url, {
    method: 'POST',
    headers: body ? { 'Content-Type': 'application/json' } : {},
    body: body ? JSON.stringify(body) : null,
  })

// Remove comment when api logout exists
//const put = url => request(url, { method: 'PUT' })

const authBackend = {
  login: (login: ?string, password: ?string) =>
    post('login', login && password ? { login, password } : null),
  logout: () => {
    // Remove comment and Promise.resolve line when api logout exists
    //return put('logout')
    return Promise.resolve()
  },
}

export default authBackend
