import { createStore, applyMiddleware, compose } from 'redux'
import thunk from 'redux-thunk'
import { createLogger } from 'redux-logger'

import rootReducer from './rootReducer'

const composeEnhancers = window.__REDUX_DEVTOOLS_EXTENSION_COMPOSE__ || compose
const devMode = 'development'

const middlewares = [thunk]

if (process.env.NODE_ENV === devMode) {
  const options = {
    collapsed: (getState, action, logEntry) => !logEntry.error,
  }
  middlewares.push(createLogger(options))
}

const store = createStore(
  rootReducer,
  composeEnhancers(applyMiddleware(...middlewares)),
)
export default store