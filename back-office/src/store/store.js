// @flow
import { createStore, applyMiddleware, compose } from 'redux'
import type { Store, Middleware } from 'redux'
import thunk from 'redux-thunk'
import { createLogger } from 'redux-logger'
import type { Options } from 'redux-logger'

import type { State, Action } from './types'
import rootReducer from './rootReducer'

const devMode = 'development'

const middlewares: Middleware<State, Action>[] = [thunk]

if (process.env.NODE_ENV === devMode) {
  const options: Options<State, Action> = {
    collapsed: (getState, action, logEntry) => !logEntry.error,
  }
  middlewares.push(createLogger(options))
}

const composeEnhancers = window.__REDUX_DEVTOOLS_EXTENSION_COMPOSE__ || compose

const store: Store<State, Action> = createStore(
  rootReducer,
  composeEnhancers(applyMiddleware(...middlewares)),
)

export default store
