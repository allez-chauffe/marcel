// @flow
import type { State as PluginsState } from '../plugins'
import type { FiltersState } from './filters'
import type { DashboardState } from '../dashboard'

export type Action = { type: string }
export type State = {
  plugins: PluginsState,
  filters: FiltersState,
  dashboard: DashboardState,
}
