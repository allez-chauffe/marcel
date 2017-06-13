// @flow
import React from 'react'
import { PluginList, PluginProps } from '../../plugins'
import './MainScreen.css'
import type { Plugin } from '../../plugins'
import { Dashboard } from '../../dashboard'

const MainScreen = ({ availablePlugins }: { availablePlugins: Plugin[] }) => (
  <div className="MainScreen">
    <div className="left-side-panel">
      <PluginList />
    </div>
    <div className="main-panel">
      <Dashboard />
    </div>
    <div className="right-side-panel">
      <PluginProps />
    </div>
  </div>
)

export default MainScreen
