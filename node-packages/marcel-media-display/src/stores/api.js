import { writable, derived } from 'svelte/store'
import { request, requestWithBody, queryParams, isUnauthorized, isNotFound } from '../utils/http'
import { config } from './config'

export const client = writable()
export const user = writable()
export const media = writable()

export const api = derived(config, ($config) => {
  if (!$config) return
  const { apiURI } = $config
  const [, base] = apiURI.match(/^(.*)\/?$/)

  const makeRequest = method => (url, options) => request(`${base}${url}`, { method, ...options })
  const makeRequestWithBody = method => (url, body, options) => requestWithBody(`${base}${url}`, body, { method, ...options })

  const get = makeRequest('GET')
  const post = makeRequestWithBody('POST')
  const put = makeRequestWithBody('PUT')
  const openWebsocket = uri => {
    return new WebSocket(`${$config.websocketURI}${uri}`)
  }

  const refreshAuth = () => post('/auth/login')
    .then(res => user.set(res))
    .catch(err => {
      if (isUnauthorized(err)) console.warn('refresh login failed')
      else throw err
    })

  const login = (login, password) => post('/auth/login', { login, password })
    .then(res => user.set(res))
    .catch(err => {
      if (isUnauthorized(err)) throw new Error('Wrong login or password')
      else throw err
    })

  const updateClient = async (loadedClient, name, mediaID) => {
    let clientToUpdate = loadedClient
    if (name && loadedClient.name !== name) clientToUpdate = { ...loadedClient, name }
    if (mediaID && loadedClient.mediaID !== mediaID) clientToUpdate = { ...loadedClient, mediaID }

    if (loadedClient === clientToUpdate) {
      client.set(loadedClient)
      return loadedClient
    } else {
      const updatedClient = await put('/clients/', clientToUpdate)
      console.info('client updated', updatedClient)
      client.set(updatedClient)
      return updateClient
    }
  }

  const createClient = async (name, mediaID) => {
    const createdClient = await post('/clients/', { name, mediaID })
    localStorage.clientID = createdClient.id
    console.info('client created', createdClient)
    client.set(createdClient)
    return createClient
  }

  const loadClient = async () => {
    const { clientID } = localStorage

    const { name, mediaID: rawMediaID } = queryParams()
    const mediaID = Number(rawMediaID)

    let loadedClient

    if (clientID) {
      loadedClient = await get(`/clients/${clientID}/`).catch(err => {
        if (isNotFound(err)) console.error(`client not found`)
        else throw err
      })
    } else {
      console.info('no client stored in local storage')
    }

    if (loadedClient) {
      console.info('client loaded', loadedClient)
      await updateClient(loadedClient, name, mediaID)
    }
    else await createClient(name, mediaID)
  }

  const connectClient = () => openWebsocket(`/clients/${localStorage.clientID}/ws`)

  const loadMedia = async (mediaID) => {
    if (!mediaID) return

    const loadedMedia = await get(`/medias/${mediaID}/`)
    console.info('media loaded', loadedMedia)
    media.set(loadedMedia)
  }

  return { loadClient, refreshAuth, login, loadMedia, connectClient }
})