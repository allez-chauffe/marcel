(() => {
  let state, props, intialized

  const nextFrame = () => new Promise(resolve => requestAnimationFrame(resolve))

  const init = PluginImplementation => {
    const instance = new PluginImplementation()
    instance.render()
    return instance
  }

  const changeProps = (newProps, prevProps) => {
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

  class Plugin {
    constructor(options = {}) {
      state = { ...options.defaultState }
      props = { ...options.defaultProps }

      addEventListener('message', async event => {
        if (event.source !== parent) return
        const message = event.data

        if (message.type === 'propsChange') {
          const { newProps, prevProps } = message.payload
          props = newProps
          await nextFrame()

          this.render()
          await nextFrame()

          if (!intialized) {
            this.pluginInitialized()
            intialized = true
            await nextFrame()
          }

          this.propsDidChange(prevProps || {}, state)
        }
      })

      parent.postMessage({ type: 'loaded' }, '*')
    }

    setState = newState => {
      const prevState = state
      state = typeof newState === 'function' ? newState(state) : { ...state, newState }

      nextFrame().then(async () => {
        this.render()
        await nextFrame()
        this.propsDidChange(props, prevState)
      })
    }

    get props() {
      return props
    }

    get state() {
      return state
    }

    pluginInitialized() {}

    propsDidChange() {}

    render() {}
  }

  window.Marcel = { Plugin, init, changeProps }
})()
