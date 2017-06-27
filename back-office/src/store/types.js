// @flow
import type { State as PluginsState } from '../plugins'
import type { FiltersState, FiltersAction } from './filters'
import type { DashboardState, DashboardAction } from '../dashboard'
import type { AuthState } from '../auth'

export type Action = FiltersAction | DashboardAction

export type Dispatch = Action => mixed

export type State = {
  plugins: PluginsState,
  filters: FiltersState,
  dashboard: DashboardState,
  auth: AuthState,
}
