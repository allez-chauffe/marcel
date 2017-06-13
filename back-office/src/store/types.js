// @flow
import type { State as PluginsState } from '../plugins'
import type { FiltersState } from './filters'

export type Action = { type: string }
export type State = {
  plugins: PluginsState,
  filters: FiltersState,
}
