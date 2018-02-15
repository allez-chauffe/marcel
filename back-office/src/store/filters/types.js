//@flow
export type FiltersAction = {
  type: 'FILTER/CHANGE_FILTER',
  payload: {
    filter: string,
    collection: string,
  },
}

export type FiltersState = {
  plugins: string,
  props: string,
  clients: string,
}
