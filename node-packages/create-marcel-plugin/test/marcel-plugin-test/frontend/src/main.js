class MarcelPluginTest extends Marcel.Plugin {
  constructor() {
    super()
    this.root = document.getElementById('root')
  }

  render() {
    const { firstName, stylesvar = {} } = this.props

    this.root.innerText = `Hello ${firstName} !`

    // stylesvar is a special property containing the global media theme.
    // You should use it to have a consistent style accross all the media.
    if (stylesvar['primary-color']) this.root.style.color = stylesvar['primary-color']
    if (stylesvar['font-family']) this.root.style.fontFamily = stylesvar['font-family']
  }
}

Marcel.init(MarcelPluginTest)

// uncomment this line to try the plugin in a browser :
// Marcel.changeProps({ firstName: 'Marcel' })
