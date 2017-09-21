function loadMedia(config) {
  console.log('Loading media ', config.client.mediaID)
  const loader = new PolymerApplicationLoader()
  const { client, backendURL, pluginURL } = config
  return loader
    .load(
      'components',
      `http://${backendURL}/medias/${client.mediaID}/`,
      `http://${pluginURL}/medias/${client.mediaID}/plugins`,
    )
    .then(media => ({ ...config, media }))
    .then(config => {
      console.log('Media loaded : ', config)
      return config
    })
}
