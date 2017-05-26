// @flow
import React from 'react'
import Input from 'react-toolbox/lib/input/Input'
import Switch from 'react-toolbox/lib/switch/Switch'
import type { PropTyped } from '../../plugins/plugins.type'

const AutoTypeField = (prop: PropTyped) => {
  if (prop.type === 'string') return <Input value={prop.value} />
  if (prop.type === 'number') return <Input type="number" value={prop.value} />
  if (prop.type === 'boolean') return <Switch checked={prop.value} />
  if (prop.type === 'json') {
    const value = JSON.stringify(prop.value, null, 2)
    return <Input multiline={true} value={value} />
  }

  return <Input value={prop.value} />
}

export default AutoTypeField
