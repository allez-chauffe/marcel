// @flow
import React from 'react'
import Tabs from 'react-toolbox/lib/tabs/Tabs'
import Tab from 'react-toolbox/lib/tabs/Tab'
import { PluginList, SubPluginProps } from '../../plugins'
import { Media, MediaConfig } from '../../components/media'
import { ClientList } from '../../clients'

import './MediaEditPage.css'

class MediaEditPage extends React.Component {
  state = { currentTab: 0 }

  onTabChange = (index: number) => this.setState({ currentTab: index })

  componentWillMount() {
    this.props.selectMedia(this.props.media.id)
  }

  render() {
    const { currentTab } = this.state
    return (
      <div className="MediaEditPage">
        <div className="left-side-panel">
          <Tabs index={currentTab} onChange={this.onTabChange}>
            <Tab label="Configuration">
              <MediaConfig />
            </Tab>
            <Tab label="Plugins">
              <PluginList />
            </Tab>
            <Tab label="Clients">
              <ClientList />
            </Tab>
          </Tabs>
        </div>
        <div className="main-panel">
          <Media />
        </div>
        <div className="right-side-panel">
          <SubPluginProps />
        </div>
      </div>
    )
  }
}

export default MediaEditPage
