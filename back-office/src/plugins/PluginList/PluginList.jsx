// @flow
import React from 'react'
import List from 'react-toolbox/lib/list/List'
import ListItem from 'react-toolbox/lib/list/ListItem'
import Input from 'react-toolbox/lib/input/Input'
import { chain } from 'lodash'

type Plugin = {
  name: string,
  elementName: string,
  icon: string,
}

class PluginList extends React.Component {
  props: { plugins: Plugin[] }
  state = { filter: '', regexp: /.*/ }

  onSearchChange = (filter: string) => {
    //Remove spaces and add 'any char' matcher between each chars.
    const regexPatern = chain(filter).split('').without(' ').join('.*').value()
    const regexp = RegExp(`.*${regexPatern}.*`, 'i')
    this.setState({ regexp, filter })
  }

  render() {
    const { plugins } = this.props
    const { filter, regexp } = this.state

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
        <Input
          label="Search plugin"
          icon="search"
          value={filter}
          onChange={this.onSearchChange}
        />
        <List selectable>
          {pluginNodes}
        </List>
      </div>
    )
  }
}

export default PluginList
