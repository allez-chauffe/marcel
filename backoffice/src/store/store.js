import { createStore, applyMiddleware, compose } from 'redux'
import thunk from 'redux-thunk'
import { createLogger } from 'redux-logger'
import { createBrowserHistory } from 'history'
import { routerMiddleware } from 'connected-react-router'

import createRootReducer from './createRootReducer'

const composeEnhancers = window.__REDUX_DEVTOOLS_EXTENSION_COMPOSE__ || compose
const devMode = 'development'

export const history = createBrowserHistory()

const middlewares = [
  thunk,
  routerMiddleware(history),
]

if (process.env.NODE_ENV === devMode) {
  const options = {
    collapsed: (_getState, _action, logEntry) => !logEntry.error,
  }
  middlewares.push(createLogger(options))
}


const store = createStore(
  createRootReducer(history),
  composeEnhancers(applyMiddleware(...middlewares)),
)

export default store
