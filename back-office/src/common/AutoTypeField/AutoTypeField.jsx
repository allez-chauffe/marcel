// @flow
import React from 'react'
import Input from 'react-toolbox/lib/input/Input'
import Switch from 'react-toolbox/lib/switch/Switch'
import type { Prop } from '../../plugins'

const AutoTypeField = ({
  value,
  onChange,
}: {
  value: Prop,
  onChange: mixed => void,
}) => {
  const { name } = value
  if (value.type === 'string')
    return <Input value={value.value} {...{ name, onChange }} />
  if (value.type === 'number')
    return <Input type="number" value={value.value} {...{ name, onChange }} />
  if (value.type === 'boolean')
    return <Switch checked={value.value} {...{ name, onChange }} />
  if (value.type === 'json') {
    return (
      <Input multiline={true} value={value.value} {...{ name, onChange }} />
    )
  }

  return <Input value={value.value} {...{ name, onChange }} />
}

export default AutoTypeField
