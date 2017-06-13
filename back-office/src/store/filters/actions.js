// @flow
import type { FiltersAction } from './types'

export const actions = {
  CHANGE_FILTER: 'FILTER/CHANGE_FILTER',
}

export const changeFilter = (collection: string) => (
  filter: string,
): FiltersAction => {
  console.log(collection, filter)
  return {
    type: actions.CHANGE_FILTER,
    payload: { filter, collection },
  }
}

export const changePluginFilter = changeFilter('plugins')
export const changePropsFilter = changeFilter('props')
