// @flow
import React from 'react'
import List from 'react-toolbox/lib/list/List'

import { SearchField } from '../../common'
import PluginListItem from './PluginListItem'
import type { Plugin } from '../type'

const PluginList = (props: {
  plugins: Plugin[],
  filter: string,
  changeFilter: string => void,
}) => {
  const { plugins, filter, changeFilter } = props

  return (
    <div>
      <SearchField
        label="Search plugin"
        value={filter}
        onChange={changeFilter}
      />
      <List selectable>
        {plugins.map(plugin => (
          <PluginListItem plugin={plugin} key={plugin.elementName} />
        ))}
      </List>
    </div>
  )
}

export default PluginList
