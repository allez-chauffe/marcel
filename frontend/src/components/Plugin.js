import React, { Component } from 'react'
import isEqual from 'lodash/isEqual'
import decorate from '../utils/decorator'

class Plugin extends Component {
  componentDidMount = () => {
    console.log('Plugin mount ')
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

  pluginProps = props => (props || this.props).plugin.frontend.props

  setIframe = iframe => {
    this.iframe = decorate(iframe)
  }

  render() {
    console.log('plugin props', this.props)
    const { instanceId } = this.props.plugin
    return (
      <iframe
        key={instanceId}
        ref={this.setIframe}
        style={this.props.style}
        className={this.props.className}
        src="iframe.html"
        title={instanceId}
        sandbox="allow-scripts allow-forms"
      />
    )
  }
}

export default Plugin
