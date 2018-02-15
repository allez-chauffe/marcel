//@flow
import { routerForBrowser } from 'redux-little-router'
import routes from './routes.js'

const router = routerForBrowser({ routes })

export const reducer = router.reducer
export const middleware = router.middleware
export const enhancer = router.enhancer
