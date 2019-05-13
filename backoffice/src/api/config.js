const config = {
  loadConfig: () =>
    fetch('/conf/config.json').then(response => {
      if (response.status !== 200) throw response
      return response.json()
    }),
}

export default config
