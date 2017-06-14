// @flow
import React from 'react'
import List from 'react-toolbox/lib/list/List'
import ListItem from 'react-toolbox/lib/list/ListItem'
import { chain } from 'lodash'

import { SearchField } from '../../common'
import type { Plugin } from '../type'

export default class PluginList extends React.Component {
  props: { plugins: Plugin[], filter: string, changeFilter: string => void }

  render() {
    const { plugins, filter, changeFilter } = this.props

    const regexPatern = chain(filter).split('').without(' ').join('.*').value()
    const regexp = RegExp(`.*${regexPatern}.*`, 'i')

    const pluginNodes = chain(plugins)
      .filter(({ name }) => regexp.test(name))
      .map(plugin => (
        <ListItem
          caption={plugin.name}
          key={plugin.elementName}
          legend={plugin.elementName}
          leftIcon={plugin.icon}
        />
      ))
      .value()

    return (
      <div>
        <SearchField
          label="Search plugin"
          value={filter}
          onChange={changeFilter}
        />
        <List selectable>
          {pluginNodes}
        </List>
      </div>
    )
  }
}
