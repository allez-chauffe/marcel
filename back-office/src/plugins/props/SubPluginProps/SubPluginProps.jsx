//@flow
import React from 'react'
import IconButton from 'react-toolbox/lib/button/IconButton'

import type { PluginInstance } from '../../../dashboard'
import PluginProps from '../PluginProps'

import './SubPluginProps.css'

export type PropsType = {
  plugin: ?PluginInstance,
  goBack: () => void,
}

const SubPluginProps = (props: PropsType) => {
  const { plugin, goBack = () => console.log('GO BACK') } = props
  return (
    <div>
      {plugin &&
        <div>
          {plugin.parent && <IconButton icon="arrow_back" onClick={goBack} />}
          <PluginProps />
        </div>}
    </div>
  )
}

export default SubPluginProps
