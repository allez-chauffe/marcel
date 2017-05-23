// @flow
import React from 'react'
import { range } from 'lodash'
import { PluginList } from '../plugins'
import './MainScreen.css'

import { Grid } from '../grid/Grid'

//TODO Remove mocked data
const availablePlugins = range(20).map(i => ({
  name: `Plugin ${i}`,
  elementName: `plugin-${i}`,
  icon: 'picture_in_picture_alt',
}))

export const MainScreen = () => (
  <div className="MainScreen">
    <div className="left-side-panel">
      <PluginList plugins={availablePlugins} />
    </div>
    <div className="main-panel">
      <Grid />
    </div>
    <div className="right-side-panel">
      right-side-panel
    </div>
  </div>
)
