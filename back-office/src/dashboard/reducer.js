//@flow
import type { Reducer } from 'redux'
import { actions } from './actions'
import { set, unset } from 'immutadot'
import uuid from 'uuid/v4'
import type { DashboardAction, DashboardState } from './type'

const intialState = {
  selectedPlugin: null,
  dashboard: {
    name: 'Dashboard',
    description: 'Some description',
    plugins: {
      'plugin-1#0': {
        name: `Plugin 1`,
        elementName: `plugin-1`,
        instanceId: 'plugin-1#0',
        icon: 'picture_in_picture_alt',
        x: 0,
        y: 0,
        columns: 2,
        rows: 3,
        props: {
          prop1: { name: 'prop1', type: 'string', value: 'hello world !' },
          prop2: { name: 'prop2', type: 'number', value: 42 },
          prop3: { name: 'prop3', type: 'boolean', value: true },
          prop4: {
            name: 'prop4',
            type: 'json',
            value: { collection: ['first', 'second'] },
          },
        },
      },
    },
  },
}

const dashboard: Reducer<DashboardState, DashboardAction> = (
  state = intialState,
  action,
) => {
  switch (action.type) {
    case actions.SELECT_PLUGIN: {
      return { ...state, selectedPlugin: action.payload.instanceId }
    }
    case actions.ADD_PLUGIN: {
      const instanceId = uuid()
      return set(state, `dashboard.plugins.${instanceId}`, {
        ...action.payload.plugin,
        x: 0,
        y: 0,
        columns: 1,
        rows: 1,
        instanceId,
      })
    }
    case actions.DELETE_PLUGIN: {
      return state.selectedPlugin
        ? unset(state, `dashboard.plugins.${state.selectedPlugin}`)
        : state
    }
    case actions.CHANGE_PROP: {
      const { instanceId, prop, value } = action.payload
      return set(
        state,
        `dashboard.plugins.${instanceId}.props.${prop.name}.value`,
        value,
      )
    }
    default:
      return state
  }
}

export default dashboard
