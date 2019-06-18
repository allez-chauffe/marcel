const config = {
  loadConfig: () =>
    fetch('config').then(response => {
      if (response.status !== 200) throw response
      return response.json()
    }),
}

export default config
