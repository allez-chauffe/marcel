import { derived } from 'svelte/store'
import * as toast from '../utils/toast'
import { api, client } from './api'

let wsConnection, clientID, errorToast

export const connection = derived([api, client], async ([$api, $client], set) => {
  const dissmissError = () => {
    if (errorToast) errorToast.close()
  }

  const displayError = (message) => {
    errorToast = toast.error(message)
  }

  const close = async () => {
    if (wsConnection) await wsConnection.close()
    set(wsConnection = null)
  }

  const open = async () => {
    set(wsConnection = await $api.connectClient())

    wsConnection.addEventListener('open', () => {
      console.info('client connected')
      dissmissError()
    })

    wsConnection.addEventListener('close', () => {
      close()
      setTimeout(open, 5000)
      dissmissError()
      displayError('Connection au backend interrompue .\nTentative de reconnexion dans 5s')
    })

    wsConnection.addEventListener('error', err => {
      console.error('client connection error. try to reconect in 5s.', err)
      close()
    })

    wsConnection.addEventListener('message', message => {
      if (message.data === 'update') {
        toast.info('Ce Media a été mis à jour')
        $api.loadClient()
      }
    })
  }

  if (!$client) return close()
  if (wsConnection && $client.id === clientID) return
  if (wsConnection) await close()
  clientID = $client.id
  return open()
})