// @flow
import React from 'react'
import List from 'react-toolbox/lib/list/List'

import { SearchField } from '../../common'
import PluginListItem from './PluginListItem'
import type { Plugin } from '../type'

type PropTypes = {
  plugins: Plugin[],
  filter: string,
  changeFilter: string => void,
}

const PluginList = (props: PropTypes) => {
  const { plugins, filter, changeFilter } = props

  return (
    <div>
      <SearchField label="Rechercher un plugin" value={filter} onChange={changeFilter} />
      <List selectable>
        {plugins.map(plugin => <PluginListItem plugin={plugin} key={plugin.eltName} />)}
      </List>
    </div>
  )
}

export default PluginList
