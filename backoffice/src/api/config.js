const config = {
  loadConfig: () =>
    fetch(window._marcelBackofficeConfigURL).then(response => {
      if (response.status !== 200) throw response
      return response.json()
    }),
}

export default config
