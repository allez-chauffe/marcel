export const pluginsSelector = state => state.plugins.list

export const pluginUpdating = (state, eltName) => state.plugins.updating[eltName]

export const addingPlugin = state => state.plugins.adding
