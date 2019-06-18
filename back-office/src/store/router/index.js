import { createReduxHistoryContext } from 'redux-first-history'
import createHistory from 'history/createBrowserHistory'

export const { createReduxHistory, routerMiddleware, routerReducer } = createReduxHistoryContext({
  history: createHistory(),
  //others options if needed
})
