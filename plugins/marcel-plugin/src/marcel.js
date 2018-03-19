;(() => {
  let state, props

  class Plugin {
    propsDidChange() {}

    render() {}

    constructor(defaults = {}) {
      state = defaults.defaultState
      props = defaults.defaultProps

      this.setState = newState => {
        state = { ...state, ...newState }
        setTimeout(() => this.render())
      }

      addEventListener('message', event => {
        if (event.source !== parent) return
        const message = event.data

        if (message.type === 'propsChange') {
          const { newProps, prevProps } = message.payload
          props = newProps
          setTimeout(() => {
            this.render()
            setTimeout(() => this.propsDidChange(prevProps || {}))
          })
        }
      })

      parent.postMessage({ type: 'loaded' }, '*')
    }

    get props() {
      return props
    }

    get state() {
      return state
    }
  }

  class Debug {
    static changeProps(newProps, prevProps) {
      dispatchEvent(
        new MessageEvent('message', {
          source: parent,
          data: {
            type: 'propsChange',
            payload: { newProps, prevProps },
          },
        }),
      )
    }
  }

  window.Marcel = { Plugin, Debug }
})()
