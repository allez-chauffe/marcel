function loadLocalConfig() {
  return fetch('conf/config.json')
    .then(res => res.json())
    .then(res => [res.backendURL || 'localhost', res.pluginURL || 'localhost'])
    .catch(() => ['localhost', 'localhost'])
    .then(ips => ({ backendURL: ips[0], pluginURL: ips[1] }))
    .then(config => {
      console.log('Local config loaded', config)
      return config
    })
}

function loadRemoteConfig(config) {
  const fetchClient = id => fetch(`http://${config.backendURL}/clients/${id}/`)

  return (
    getClientId(config.backendURL)
      .then(fetchClient)
      //Check if the client has been deleted. Renewing id if nout found.
      .then(res => {
        if (res.status === 404) {
          localStorage.clear()
          return getClientId(config.backendURL).then(fetchClient)
        }
        return res
      })
      //Checking for errors
      .then(res => {
        if (res.status != 200) throw res
        return res.json()
      })
      //Checking if a media ID has specified in query params. If any, overrides the server configuration
      .then(client => {
        let { mediaID } = window.location.queryParams
        mediaID = parseInt(mediaID)

        if (mediaID && mediaID != client.mediaID) {
          const newClient = { ...client, mediaID }
          return fetch(`http://${config.backendURL}/clients/`, {
            method: 'PUT',
            body: JSON.stringify(newClient),
          }).then(res => {
            if (res.status !== 204) throw res
            return newClient
          })
        }

        return client
      })
      //Merge configration
      .then(client => ({ ...config, client }))
      .then(config => {
        console.log('Remote config loaded', config)
        return config
      })
  )
}

function getClientId(backendURL) {
  const { clientID } = localStorage

  if (clientID) {
    console.log('Client id found in local storage :', clientID)
    return Promise.resolve(clientID)
  }

  return fetch(`http://${backendURL}/clients/`, {
    method: 'POST',
    body: JSON.stringify({
      name: window.location.queryParams.name,
      mediaID: window.location.queryParams.mediaID,
    }),
  })
    .then(res => {
      if (res.status !== 200) throw res
      return res.json()
    })
    .then(client => {
      localStorage.clientID = client.id
      console.log('New client id created by server : ', client.id)
      return client.id
    })
}
