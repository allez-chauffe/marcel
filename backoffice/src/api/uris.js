export default {
  load: () =>
    fetch('uris').then(response => {
      if (response.status !== 200) {
        if (response.status === 404) {
          console.warn('No URIs config available') // eslint-disable-line no-console
          return { API: '/api/' }
        }
        throw response
      }
      return response.json()
    }),
}
