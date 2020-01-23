import { createMachine, state, transition, invoke, immediate, guard, reduce, action } from 'robot3'
import * as toast from '../utils/toast'
import { storeError, transitionWithData } from './utils'
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
  transitionWithData('done', 'loggedIn', 'user'),
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
    initial: state(
      transitionWithData('configChanged', 'loadConfig', 'config')
    ),

    // Config
    loadConfig: state(
      immediate('loadingConfig', guard(ctx => !ctx.config.apiURI || !ctx.config.websocketURI)),
      immediate('refreshingLogin')
    ),
    loadingConfig: invoke(loadConfig,
      transitionWithData('done', 'refreshingLogin', 'config'),
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
      transitionWithData('done', 'connectingClient', 'client', action(({ client }) => localStorage.clientID = client.id)),
    ),
    loadingClient: invokeHttpWithError(
      loadClient,
      transitionWithData('done', 'clientLoaded', 'client'),
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
    updatingClient: invokeHttp(
      updateClient,
      transitionWithData('done', 'connectingClient', 'client')
    ),
    reloadClient: invokeHttpWithError(
      loadClient,
      transitionWithData('done', 'loadMedia', 'client'),
    ),
    connectingClient: invokeWithError(
      connectClient,
      transitionWithData('done', 'loadMedia', 'connection')
    ),


    // Media
    loadMedia: state(
      immediate('noMedia', guard(({ client }) => !client.mediaID)),
      immediate('loadingMedia')
    ),
    loadingMedia: invokeHttpWithError(
      loadMedia,
      transitionWithData('done', 'mediaLoaded', 'media')
    ),
    mediaLoaded: connectedState('mediaLoaded'),
    noMedia: connectedState('noMedia'),

    // Errors
    fatalError: state(
      transitionWithData('configChanged', 'loadConfig', 'config')
    )
  },
  () => initialContext
)

export default machine