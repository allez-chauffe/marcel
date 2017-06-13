// @flow
import React from 'react'
import List from 'react-toolbox/lib/list/List'
import ListItem from 'react-toolbox/lib/list/ListItem'
import IconButton from 'react-toolbox/lib/button/IconButton'

import { SearchField } from '../../common'
import type { Plugin } from '../type'

const AddButton = ({ onClick }) => (
  <IconButton icon="add" primary onClick={onClick} />
)

export default class PluginList extends React.Component {
  props: {
    plugins: Plugin[],
    filter: string,
    changeFilter: string => void,
    addPlugin: Plugin => void,
  }

  render() {
    const { plugins, filter, changeFilter, addPlugin } = this.props

    return (
      <div>
        <SearchField
          label="Search plugin"
          value={filter}
          onChange={changeFilter}
        />
        <List selectable>
          {plugins.map(plugin => {
            const { name, elementName, icon } = plugin
            return (
              <ListItem
                caption={name}
                ripple={false}
                key={elementName}
                legend={elementName}
                leftIcon={icon}
                rightIcon={<AddButton onClick={() => addPlugin(plugin)} />}
              />
            )
          })}
        </List>
      </div>
    )
  }
}
