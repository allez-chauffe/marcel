// @flow
import React from 'react'

import { SearchField } from '../../../common'
import PluginProp from '../PluginProp'
import type { Plugin } from '../../plugins.type'

import './PluginProps.css'

class PluginProps extends React.Component {
  props: { plugin: Plugin }
  state: { filter: string } = { filter: '' }

  onFilterChange = (filter: string) => this.setState({ filter })

  render() {
    const { plugin } = this.props
    const { name, props } = plugin

    return (
      <div className="PluginProps">
        <h2>{name}</h2>
        <SearchField
          label="Search Prop"
          value={this.state.filter}
          onChange={this.onFilterChange}
        />
        {props.map(p => <PluginProp prop={p} key={p.name} />)}
      </div>
    )
  }
}

export default PluginProps
