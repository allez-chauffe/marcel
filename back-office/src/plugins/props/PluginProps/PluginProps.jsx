// @flow
import React from 'react'
import { values } from 'lodash'
import Button from 'react-toolbox/lib/button/Button'

import { SearchField } from '../../../common'
import PluginProp from '../PluginProp'
import type { PluginInstance } from '../../../dashboard'

import './PluginProps.css'

class PluginProps extends React.Component {
  props: {
    plugin: ?PluginInstance,
    filter: string,
    changeFilter: string => void,
    deletePlugin: PluginInstance => void,
  }

  deletePlugin = () => {
    if (this.props.plugin) this.props.deletePlugin(this.props.plugin)
  }

  render() {
    const { plugin, filter, changeFilter } = this.props

    if (!plugin) return <div className="PluginsProps" />

    const { name, props: pluginProps, x, y, cols, rows } = plugin

    return (
      <div className="PluginProps">
        <h2>
          {name}
        </h2>
        <p>{`(x: ${x}, y: ${y}, columns: ${cols}, rows: ${rows})`}</p>

        <SearchField
          label="Search Prop"
          value={filter}
          onChange={changeFilter}
        />

        {values(pluginProps).map(p =>
          <PluginProp plugin={plugin} prop={p} key={p.name} />,
        )}

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
