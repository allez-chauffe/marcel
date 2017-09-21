//@flow
import React from 'react'
import Button from 'react-toolbox/lib/button/Button'

import type { Dashboard } from '../../dashboard/type'

import './OpenButton.css'

class OpenButton extends React.Component {
  props: {
    dashboard: Dashboard,
    frontendURI: string,
  }

  open = (event: Event) => {
    const { frontendURI, dashboard } = this.props
    event.stopPropagation()
    window.open(`${frontendURI}?mediaID=${dashboard.id}`)
    window.focus()
  }

  render = () => <Button label="Ouvrir" icon="exit_to_app" onClick={this.open} />
}

export default OpenButton
