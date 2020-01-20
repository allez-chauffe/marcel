const config = {
  loadConfig: () =>
    fetch('/uris').then(response => {
      if (response.status !== 200) {
        if (response.status === 404) {
          console.warn('No config available') // eslint-disable-line no-console
          return { API: '/api/' }
        }
        throw response
      }
      return response.json()
    }),
}

export default config
