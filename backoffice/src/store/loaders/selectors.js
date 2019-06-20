const isLoading = ressource => state => state.loaders[ressource]

export const isConfigLoading = isLoading('config')

export const isPluginsLoading = isLoading('plugins')

export const isDashboardsLoading = isLoading('dashboards')

export const isLoadingInitData = isLoading('initial')

export const isClientsLoading = isLoading('clients')

export const isUsersLoading = isLoading('users')
