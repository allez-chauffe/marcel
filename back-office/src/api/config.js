//@flow
import type { Config } from '../store'

const config = {
  loadConfig: (): Promise<Config> =>
    fetch('/conf/config.json').then(response => {
      if (response.status !== 200) throw response
      return response.json()
    }),
}

export default config
