// @flow
import type { State as PluginsState } from '../plugins'
import type { FiltersState, FiltersAction } from './filters'
import type { DashboardState, DashboardAction } from '../dashboard'
import type { AuthState, AuthAction } from '../auth'
import type { LoadersState, LoadersAction } from './loaders'
import type { ClientState, ClientAction } from '../clients'

export type Config = {
  backendURI: string,
  authURI: string,
  frontendURI: string,
}

export type Action = FiltersAction | DashboardAction | LoadersAction | AuthAction | ClientAction

export type Dispatch = Action => mixed

export type State = {
  plugins: PluginsState,
  filters: FiltersState,
  dashboard: DashboardState,
  clients: ClientState,
  auth: AuthState,
  loaders: LoadersState,
  config: Config,
}
