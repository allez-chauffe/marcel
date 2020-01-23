import { createMachine, state, transition, invoke, immediate, guard, reduce, interpret } from 'robot3'
import { storeError, transitionWithData } from './utils'
import { invokeHttp, isUnauthorized } from './http'
import { login, refreshLogin, loadConfig, createClient } from './api'
import { writable } from 'svelte/store'
import 'robot3/debug'

const fatalErrorTransition = () => transition('error', 'fatalError', reduce(storeError()))

const invokeHttpWithError = (...args) => invokeHttp(
  ...args,
  fatalErrorTransition()
)

const loginState = (loginFunction, unauthorizedReducer) => invokeHttpWithError(
  loginFunction,
  transitionWithData('done', 'loggedIn', 'user'),
  transition('error', 'loggedOut', guard(isUnauthorized), unauthorizedReducer),
)

const initialContext = {}

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
      immediate('creatingCient', guard(() => !localStorage.clientID)),
      immediate('loadingClient', reduce(ctx => ({ clientID: localStorage.clientID, ...ctx }))),
    ),

    // Client
    creatingClient: invokeHttpWithError(
      createClient,
      transitionWithData('done', 'connectingClient', 'client'),
    ),

    // Errors
    fatalError: state(
      transitionWithData('configChanged', 'loadConfig', 'config')
    )
  },
  () => initialContext
)


const { subscribe, set } = writable({ state: machine.current, context: initialContext, machine })

const service = interpret(
  machine,
  ({ machine, context }) => {
    // eslint-disable-next-line no-console
    console.debug('changed state', machine.current, context)
    set({ state: machine.current, context, machine })
  }
)

const send = (...args) => {
  // eslint-disable-next-line no-console
  console.debug('event sent', ...args)
  return service.send(...args)
}

export default { subscribe, send }