import { request } from '../../utils/http'
import * as http from './http'

export const loadConfig = () => request('./config.json')

export const refreshLogin = () => http.post('/auth/login')

export const login = (ctx, { login, password }) => http.post('/auth/login', { login, password })

export const createClient = (ctx, { name, mediaID }) => http.post('/clients/', { name, mediaID })