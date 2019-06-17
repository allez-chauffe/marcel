import React from 'react'
import './PluginProp.css'
import { AutoTypeField } from '../../../common'

const PluginProp = props => {
  const { prop, changeProp } = props
  return (
    <div className="PluginProp">
      <div className="propName">{prop.name}</div>
      <div className="propValue">
        <AutoTypeField value={prop} onChange={changeProp} />
      </div>
    </div>
  )
}

export default PluginProp
