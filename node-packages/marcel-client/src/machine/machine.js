import { createMachine, state, transition, invoke, immediate, guard, reduce, action, state as final } from 'robot3'
import * as toast from '../utils/toast'
import { storeError, storeDataWithoutError } from './utils'
import { invokeHttp, isUnauthorized } from './http'
import { login, refreshLogin, loadConfig, createClient, loadClient, updateClient, connectClient, loadMedia } from './api'
import 'robot3/debug'
import { queryParams } from '../utils/http'

const fatalErrorTransition = (...args) => transition('error', 'fatalError', reduce(storeError()), ...args)

const invokeWithError = (...args) => invoke(
  ...args,
  fatalErrorTransition()
)

const invokeHttpWithError = (...args) => invokeHttp(
  ...args,
  fatalErrorTransition()
)

const loginState = (loginFunction, unauthorizedReducer) => invokeHttpWithError(
  loginFunction,
  transition('done', 'loggedIn', storeDataWithoutError('user')),
  transition('error', 'loggedOut', guard(isUnauthorized), unauthorizedReducer),
)

const connectedState = (errorState, ...transitions) => state(
  transition('clientUpdated', 'reloadClient'),
  transition('clientConnectionError', errorState, action(() => toast.error('Une erreur de connection avec le serveur est survenue !'))),
  transition('clientConnectionClosed', 'connectingClient', action(() => toast.warning('La connection au serveur est interrompue'))),
  ...transitions,
)

export const initialContext = {}

const machine = createMachine(
  {
    // Config
    loadingConfig: invoke(loadConfig,
      transition('done', 'refreshingLogin', storeDataWithoutError('config')),
      transition('error', 'refreshingLogin', reduce((ctx, { error }) => {
        console.error('enable to fetch config, using default', error.stack)
        return { config: { apiURI: '/api/' } }
      })),
    ),

    // User
    refreshingLogin: loginState(refreshLogin),
    loggingIn: loginState(login, reduce(storeError("Login ou mot de passe incorrect"))),
    loggedOut: state(
      transition('login', 'loggingIn')
    ),
    loggedIn: state(
      immediate('creatingClient', guard(() => !localStorage.clientID)),
      immediate('loadingClient'),
    ),

    // Client
    creatingClient: invokeHttpWithError(
      createClient,
      transition('done', 'connectingClient',
        storeDataWithoutError('client'),
        action(({ client }) => localStorage.clientID = client.id)
      ),
    ),
    loadingClient: invokeHttpWithError(
      loadClient,
      transition('done', 'clientLoaded', storeDataWithoutError('client')),
    ),
    clientLoaded: state(
      immediate('updatingClient',
        guard(({ client }) => {
          const { name, mediaID } = queryParams()
          return name && client.name !== name || mediaID && mediaID !== client.mediaID.toString()
        })
      ),
      immediate('connectingClient')
    ),
    updatingClient: invokeHttpWithError(
      updateClient,
      transition('done', 'connectingClient', storeDataWithoutError('client'))
    ),
    reloadClient: invokeHttpWithError(
      loadClient,
      transition('done', 'loadMedia', storeDataWithoutError('client')),
    ),
    connectingClient: invokeWithError(
      connectClient,
      transition('done', 'loadMedia', storeDataWithoutError('connection'))
    ),


    // Media
    loadMedia: state(
      immediate('noMedia', guard(({ client }) => !client.mediaID)),
      immediate('loadingMedia')
    ),
    loadingMedia: invokeHttpWithError(
      loadMedia,
      transition('done', 'mediaLoaded', storeDataWithoutError('media'))
    ),
    mediaLoaded: connectedState('mediaLoaded'),
    noMedia: connectedState('noMedia'),

    // Errors
    fatalError: final()
  },
  () => initialContext
)

export default machine
