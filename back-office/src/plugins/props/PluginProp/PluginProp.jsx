// @flow
import React from 'react'
import './PluginProp.css'
import type { Prop } from '../../plugins.type'
import { AutoTypeField } from '../../../common'

const PluginProp = (props: { prop: Prop }) => {
  const { type, name, value } = props.prop
  return (
    <div className="PluginProp">
      <div className="propName">{name}</div>
      <div className="propValue">
        <AutoTypeField type={type} value={value} />
      </div>
    </div>
  )
}

export default PluginProp
