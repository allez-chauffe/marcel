// @flow
import React from 'react'
import { values } from 'lodash'

import { SearchField } from '../../../common'
import PluginProp from '../PluginProp'
import type { PluginInstance } from '../../../dashboard'

import './PluginProps.css'

const PluginProps = (props: {
  plugin?: PluginInstance,
  filter: string,
  changeFilter: string => void,
}) => {
  const { plugin, filter, changeFilter } = props

  if (!plugin) return <div className="PluginsProps" />

  const { name, props: pluginProps } = plugin

  return (
    <div className="PluginProps">
      <h2>{name}</h2>
      <SearchField label="Search Prop" value={filter} onChange={changeFilter} />
      {values(pluginProps).map(p => (
        <PluginProp plugin={plugin} prop={p} key={p.name} />
      ))}
    </div>
  )
}

export default PluginProps
