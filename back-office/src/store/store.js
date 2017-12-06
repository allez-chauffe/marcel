// @flow
import { createStore, applyMiddleware, compose } from 'redux'
import type { Store, Middleware } from 'redux'
import thunk from 'redux-thunk'
import { createLogger } from 'redux-logger'
import type { Options } from 'redux-logger'

import type { State, Action } from './types'
import * as router from './router'
import rootReducer from './rootReducer'

const composeEnhancers = window.__REDUX_DEVTOOLS_EXTENSION_COMPOSE__ || compose
const devMode = 'development'

const middlewares: Middleware<State, Action>[] = [thunk]

if (process.env.NODE_ENV === devMode) {
  const options: Options<State, Action> = {
    collapsed: (getState, action, logEntry) => !logEntry.error,
  }
  middlewares.push(createLogger(options))
}

const store: Store<State, Action> = createStore(
  rootReducer,
  composeEnhancers(router.enhancer, applyMiddleware(router.middleware, ...middlewares)),
)

export default store
