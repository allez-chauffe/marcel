// @flow
import React from 'react'
import Input from 'react-toolbox/lib/input/Input'
import Switch from 'react-toolbox/lib/switch/Switch'
import { PluginListProp } from '../../plugins'
import type { Prop } from '../../plugins'

export type PropsType = {
  value: Prop,
  onChange: mixed => void,
}

const AutoTypeField = (props: PropsType) => {
  const { value, onChange } = props
  const { name } = value
  if (value.type === 'string')
    return <Input value={value.value} name={name} onChange={onChange} />
  if (value.type === 'number')
    return (
      <Input
        type="number"
        value={value.value}
        name={name}
        onChange={onChange}
      />
    )
  if (value.type === 'boolean')
    return <Switch checked={value.value} name={name} onChange={onChange} />
  if (value.type === 'json') {
    return (
      <Input
        multiline={true}
        value={value.value}
        name={name}
        onChange={onChange}
      />
    )
  }
  if (value.type === 'pluginList') {
    return (
      <PluginListProp name={name} value={value.value} onChange={onChange} />
    )
  }

  return <Input value={value.value} name={name} onChange={onChange} />
}

export default AutoTypeField
