// @flow
import React from 'react'
import { range } from 'lodash'
import { PluginList, PluginProps } from '../plugins'
import './MainScreen.css'

//TODO Remove mocked data
const availablePlugins = range(20).map(i => ({
  name: `Plugin ${i}`,
  elementName: `plugin-${i}`,
  icon: 'picture_in_picture_alt',
  props: [
    { name: 'prop1', type: 'string', value: 'coucou' },
    { name: 'prop2', type: 'number', value: 5 },
    { name: 'prop3', type: 'boolean', value: true },
    { name: 'prop4', type: 'json', value: [{ machin: 1 }] },
  ],
}))

export const MainScreen = () => (
  <div className="MainScreen">
    <div className="left-side-panel">
      <PluginList plugins={availablePlugins} />
    </div>
    <div className="main-panel">
      main-panel
    </div>
    <div className="right-side-panel">
      <PluginProps plugin={availablePlugins[0]} />
    </div>
  </div>
)
