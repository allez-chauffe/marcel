// @flow
import type { State as PluginsState } from '../plugins'
import type { FiltersState, FiltersAction } from './filters'
import type { DashboardState, DashboardAction } from '../dashboard'
import type { AuthState, AuthAction } from '../auth'
import type { LoadersState, LoadersAction } from './loaders'

export type Config = {
  backendURI: string,
  frontendURI: string,
}

export type Action = FiltersAction | DashboardAction | LoadersAction | AuthAction

export type Dispatch = Action => mixed

export type State = {
  plugins: PluginsState,
  filters: FiltersState,
  dashboard: DashboardState,
  auth: AuthState,
  loaders: LoadersState,
  config: Config,
}
