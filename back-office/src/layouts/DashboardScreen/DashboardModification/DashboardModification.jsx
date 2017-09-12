// @flow
import React from 'react'
import Tabs from 'react-toolbox/lib/tabs/Tabs'
import Tab from 'react-toolbox/lib/tabs/Tab'
import { PluginList, SubPluginProps } from '../../../plugins'
import { Dashboard, DashboardConfig } from '../../../dashboard'

import './DashboardModification.css'

class DashboardModification extends React.Component {
  state = { currentTab: 0 }

  onTabChange = (index: number) => this.setState({ currentTab: index })

  render() {
    const { currentTab } = this.state
    return (
      <div className="DashboardModification">
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
          <SubPluginProps />
        </div>
      </div>
    )
  }
}

export default DashboardModification
