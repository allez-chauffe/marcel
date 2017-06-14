// @flow
import React from 'react'
import './PluginProp.css'
import type { Prop } from '../../type'
import { AutoTypeField } from '../../../common'

const PluginProp = ({ prop }: { prop: Prop }) => (
  <div className="PluginProp">
    <div className="propName">{prop.name}</div>
    <div className="propValue">
      <AutoTypeField value={prop} />
    </div>
  </div>
)

export default PluginProp
