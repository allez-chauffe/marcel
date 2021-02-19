import React from 'react'
import Button from 'react-toolbox/lib/button/Button'

import './OpenButton.css'

class OpenButton extends React.Component {
  open = event => {
    const { Frontend, dashboard } = this.props
    event.stopPropagation()
    window.open(`${Frontend}?mediaID=${dashboard.id}`)
    window.focus()
  }

  render = () => <Button label="Ouvrir" icon="exit_to_app" onClick={this.open} />
}

export default OpenButton
