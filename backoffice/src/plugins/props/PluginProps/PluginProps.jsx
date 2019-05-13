import React, { Component } from 'react'
import { values } from 'lodash'
import Button from 'react-toolbox/lib/button/Button'

import { SearchField } from '../../../common'
import PluginProp from '../PluginProp'

import './PluginProps.css'

class PluginProps extends Component {
  deletePlugin = () => {
    if (this.props.plugin) this.props.deletePlugin(this.props.plugin)
  }

  render() {
    const { plugin, filter, changeFilter } = this.props

    if (!plugin) return <div className="PluginsProps" />

    const { props: pluginProps } = plugin

    return (
      <div className="PluginProps">
        <SearchField label="Search Prop" value={filter} onChange={changeFilter} />

        {values(pluginProps).map(p => (
          <PluginProp plugin={plugin} prop={p} key={p.name} />
        ))}

        <Button
          icon="delete"
          label="Supprimer"
          raised
          primary
          className="delete"
          onClick={this.deletePlugin}
        />
      </div>
    )
  }
}

export default PluginProps
