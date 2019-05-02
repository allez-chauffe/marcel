import React from 'react'

const marcel = Component => {
  class MarcelHOC extends React.Component {
    state = {
      props: {},
    }

    componentDidMount() {
      this.eventListener = window.addEventListener('message', this.onMessage)
      window.parent.postMessage({ type: 'loaded' }, '*')
    }

    componentWillUnmount() {
      window.removeEventListener('message', this.eventListener)
    }

    onMessage = event => {
      if (event.source !== window.parent) return
      const { type, payload } = event.data

      switch (type) {
        case 'propsChange':
          console.log('Message received: ', payload)
          return this.setState({ props: payload.newProps })
        default:
      }
    }

    render = () => {
      return <Component {...this.props} {...this.state.props} />
    }
  }

  return MarcelHOC
}

export default marcel
