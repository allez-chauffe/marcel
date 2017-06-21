// @flow
import React from 'react'
import Tabs from 'react-toolbox/lib/tabs/Tabs'
import Tab from 'react-toolbox/lib/tabs/Tab'
import { PluginList, PluginProps } from '../../plugins'
import { Dashboard, DashboardConfig } from '../../dashboard'

import './MainScreen.css'

class MainScreen extends React.Component {
  state = { currentTab: 0 }

  onTabChange = (index: number) => this.setState({ currentTab: index })

  render() {
    const { currentTab } = this.state
    return (
      <div className="MainScreen">
        <div className="left-side-panel">
          <Tabs index={currentTab} onChange={this.onTabChange}>
            <Tab label="Configuration">
              <DashboardConfig />
            </Tab>
            <Tab label="Plugins">
              <PluginList />
            </Tab>
          </Tabs>
        </div>
        <div className="main-panel">
          <Dashboard />
        </div>
        <div className="right-side-panel">
          <PluginProps />
        </div>
      </div>
    )
  }
}

export default MainScreen
