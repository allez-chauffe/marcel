// @flow
import { createStore, applyMiddleware } from 'redux'
import type { Store, Middleware } from 'redux'
import thunk from 'redux-thunk'
import { createLogger } from 'redux-logger'
import type { Options } from 'redux-logger'

import type { State, Action } from './types'
import rootReducer from './rootReducer'

const middlewares: Middleware<State, Action>[] = [thunk]

if (process.env.NODE_ENV === `development`) {
  const options: Options<State, Action> = {
    collapsed: (getState, action, logEntry) => !logEntry.error,
  }
  middlewares.push(createLogger(options))
}

const store: Store<State, Action> = createStore(
  rootReducer,
  applyMiddleware(...middlewares),
)

export default store
