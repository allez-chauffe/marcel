//@flow
import React from 'react'
import Dropdown from 'react-toolbox/lib/dropdown/Dropdown'
import type { Plugin } from '../../../../plugins'

export type PropsType = {
  plugins: Plugin[],
  addPlugin: Plugin => void,
}

const AddPlugin = (props: PropsType) => {
  const { plugins, addPlugin } = props
  const source = [
    { value: -1, label: 'Ajouter un plugin' },
    ...plugins.map(plugin => ({ value: plugin, label: plugin.name })),
  ]

  return <Dropdown source={source} value={-1} onChange={addPlugin} />
}

export default AddPlugin
