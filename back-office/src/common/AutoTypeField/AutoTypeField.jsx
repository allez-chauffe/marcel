// @flow
import React from 'react'
import Input from 'react-toolbox/lib/input/Input'
import Switch from 'react-toolbox/lib/switch/Switch'
import type { Prop } from '../../plugins'

const AutoTypeField = ({ value }: { value: Prop }) => {
  if (value.type === 'string')
    return <Input value={value.value} name={value.name} />
  if (value.type === 'number')
    return <Input type="number" value={value.value} name={value.name} />
  if (value.type === 'boolean')
    return <Switch checked={value.value} name={value.name} />
  if (value.type === 'json') {
    const stringValue = JSON.stringify(value.value, null, 2)
    return <Input multiline={true} value={stringValue} name={value.name} />
  }

  return <Input value={value.value} />
}

export default AutoTypeField
