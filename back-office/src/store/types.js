// @flow
import type { State as PluginsState } from '../plugins'

export type Action = { type: string }
export type State = {
  plugins: PluginsState,
}
