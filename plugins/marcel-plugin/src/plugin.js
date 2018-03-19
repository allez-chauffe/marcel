class Text extends Marcel.Plugin {
  constructor() {
    super({
      defaultProps: {
        text: '',
        stylesvar: {},
      },
    })

    this.content = document.getElementById('content')
  }

  render() {
    const { text, stylesvar } = this.props

    this.content.innerText = text

    if (stylesvar['primary-color']) this.p.style.color = stylesvar['primary-color']
    if (stylesvar['font-family']) this.p.style.fontFamily = stylesvar['font-family']
  }
}

const instance = new Text()
