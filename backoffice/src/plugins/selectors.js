export const pluginsSelector = state => state.plugins.list

export const pluginUpdating = (eltName, state) => state.plugins.updating[eltName]

export const addingPlugin = state => state.plugins.adding
