import React, { Component } from 'react'
import isEqual from 'lodash/isEqual'
import decorate from '../utils/decorator'

class Plugin extends Component {
  componentDidMount = () => {
    console.log('Plugin mount ')
    // this.iframe.window.parent
    this.iframe.addMessageListener(message => {
      console.log(`Message from ${this.props.plugin.instanceId}`, message)
      if (message.type === 'loaded') this.emitChangeProps(this.pluginProps(), null)
    })
  }

  shouldComponentUpdate = nextProps => !isEqual(this.props, nextProps)

  componentDidUpdate = previousProps =>
    this.emitChangeProps(this.pluginProps(), previousProps.plugin.frontend.props)

  emitChangeProps = (newProps, prevProps) =>
    this.iframe.postMessage({ type: 'propsChange', payload: { newProps, prevProps } })

  pluginProps = () => ({ ...this.props.plugin.frontend.props, stylesvar: this.props.stylesvar })

  setIframe = iframe => {
    this.iframe = decorate(iframe)
  }

  render() {
    const { pluginsURL, style, className, plugin: { instanceId, eltName, name } } = this.props
    return (
      <iframe
        key={instanceId}
        ref={this.setIframe}
        style={style}
        className={className}
        src={`${pluginsURL}/${eltName}/${instanceId}/`}
        title={name}
        // sandbox="allow-scripts allow-forms"
      />
    )
  }
}

export default Plugin
