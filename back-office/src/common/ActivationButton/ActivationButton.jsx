//@flow
import React from 'react'
import IconButton from 'react-toolbox/lib/button/IconButton'

import './ActivationButton.css'

class ActivationButton extends React.Component {
  props: {
    isActive: boolean,
    activate: () => void,
    deactivate: () => void,
  }

  onClick = (event: Event) => {
    event.stopPropagation()
    if (this.props.isActive) this.props.deactivate()
    else this.props.activate()
  }

  render() {
    if (!this.props.isWritable) return null

    const className = `ActivationButton ${this.props.isActive ? 'active' : 'not-active'}`

    return (
      <div className={className}>
        <IconButton icon="power_settings_new" onClick={this.onClick} />
      </div>
    )
  }
}

export default ActivationButton
