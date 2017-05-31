// @flow
import React from 'react'
import { PluginList, PluginProps } from '../../plugins'
import './MainScreen.css'
import type { Plugin } from '../../plugins/plugins.type'
import { Dashboard } from '../../grid'

const MainScreen = ({ availablePlugins }: { availablePlugins: Plugin[] }) => (
  <div className="MainScreen">
    <div className="left-side-panel">
      <PluginList plugins={availablePlugins} />
    </div>
    <div className="main-panel">
      <Dashboard />
    </div>
    <div className="right-side-panel">
      <PluginProps plugin={availablePlugins[0]} />
    </div>
  </div>
)

export default MainScreen
