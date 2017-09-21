function openWebsocketConnection(config) {
  const { backendURL, client, media } = config
  if (!client.mediaID) {
    console.log('No media associated')
    showError(
      `Ce client (${client.name}) n'est associé à aucun Media.\nVous pouvez associer un Media dans le backoffice`,
    )
  } else if (!media || !media.isactive) {
    console.log('Media not activated')
    showError(
      `Ce Media (${media.name} n'est pas acitvé.\nVous pouvez activer ce Media dans le backoffice.`,
    )
  }

  const conn = new WebSocket(`ws://${backendURL}/clients/${client.id}/ws`)

  conn.onmessage = event => {
    console.log('Message received from server : ', event)
    const { data } = event

    if (data === 'update') {
      showError('Ce Media a été mis à jour.\n\n Il va être rechargé dans quelques secondes...')
      setTimeout(() => window.location.reload(), 5000)
    }
  }

  conn.onopen = () => {
    console.log('Connection with server established')
    if (client.mediaID && media && media.isactive) hideError()
  }

  conn.onclose = event => {
    setTimeout(
      () => loadRemoteConfig(config).then(newConfig => openWebsocketConnection(newConfig)),
      5000,
    )
    console.log('Connection with backend lost')
    showError('La connection avec le backend a été interrompue.\nTentative de reconnexion...')
  }

  window.onbeforeunload = () => {
    conn.close()
  }
}
