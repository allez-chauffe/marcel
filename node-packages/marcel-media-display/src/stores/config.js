import { writable } from 'svelte/store'
import { request } from '../utils/http'
import { withoutNil } from '../utils/object'
import { debounce } from '../utils/function'

export const config = writable()

export const loadConfig = debounce(async (givenConfig) => {
  let loadedConfig
  try {
    const response = await request('./config')
    loadedConfig = await response.json()
  } catch (err) {
    console.warn('An error occured while loading config', err, err.response)
    loadedConfig = { apiURI: '/api/', ...givenConfig }
  }

  const completeConfig = { ...loadedConfig, ...withoutNil(givenConfig) }
  console.info('config loaded', completeConfig)
  config.set(completeConfig)
})