import React from 'react'

import { PluginCard, AddPlugin } from '../../plugins'

import './PluginsScreen.css'

const PluginsScreen = ({ plugins }) => (
  <div className="PluginsScreen">
    <AddPlugin />
    <div className="CardGrid">
      {plugins.map(plugin => (
        <PluginCard key={plugin.eltName} plugin={plugin} />
      ))}
    </div>
  </div>
)

export default PluginsScreen
