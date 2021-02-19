import React from 'react'
import IconButton from 'react-toolbox/lib/button/IconButton'

import PluginProps from '../PluginProps'

import './SubPluginProps.css'

const SubPluginProps = props => {
  const { plugin, goBack } = props
  if (!plugin) return <div />

  const { x, y, cols, rows, name, parent } = plugin
  return (
    <div className="SubPluginProps">
      <div>
        <h2>
          {parent && <IconButton icon="arrow_back" onClick={goBack} className="back" />}
          {name}
        </h2>
        <p>{`(x: ${x}, y: ${y}, columns: ${cols}, rows: ${rows})`}</p>
      </div>
      <PluginProps />
    </div>
  )
}

export default SubPluginProps
