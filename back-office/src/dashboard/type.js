//@flow

export type SelectPluginAction = {
  type: string,
  payload: {
    elementName: string,
  },
}

export type DashboardAction = SelectPluginAction

export type DashboardState = {
  selectedPlugin: string | null,
}
