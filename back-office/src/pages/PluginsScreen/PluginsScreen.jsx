import React from 'react'
import {PluginCard} from '../../plugins'

import './PluginsScreen.css'

const PluginsScreen = ({ plugins }) => {
  return (
    <div className="PluginsScreen CardGrid">
      {plugins.map(plugin => (
        <PluginCard key={plugin.eltName} plugin={plugin} />
      ))}
    </div>
  )
}

export default PluginsScreen
