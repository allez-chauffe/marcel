// @flow
import React from 'react'
import { PluginList } from '../plugins'
import './MainScreen.css'
import type { Plugin } from '../plugins/plugins.types'
import { Dashboard } from '../grid'

const MainScreen = ({ availablePlugins }: { availablePlugins: Plugin[] }) => (
  <div className="MainScreen">
    <div className="left-side-panel">
      <PluginList plugins={availablePlugins} />
    </div>
    <div className="main-panel">
      <Dashboard />
    </div>
    <div className="right-side-panel">
      right-side-panel
    </div>
  </div>
)

export default MainScreen
