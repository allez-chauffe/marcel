import React, { Component } from 'react'
import { ChromePicker } from 'react-color'

import './ColorPicker.css'

class ColorPicker extends Component {
  state = {
    displayed: false,
    oldColor: '',
  }

  onKeypressed = event => {
    // Enter pressed
    if (event.keyCode === 13) this.close()

    // Excape pressed
    if (event.keyCode === 27) {
      this.onChange({ hex: this.state.oldColor })
      this.close()
    }
  }

  open = () => {
    document.addEventListener('keyup', this.onKeypressed)
    this.setState({ oldColor: this.props.value, displayed: true })
  }

  close = () => {
    document.removeEventListener('keyup', this.onKeypressed)
    this.setState({ displayed: false })
  }

  onChange = newColor => this.props.onChange(newColor.hex)

  render() {
    const { label, value } = this.props
    const { displayed } = this.state
    const { onChange, close, open } = this
    return (
      <div className="ColorPicker">
        <div onClick={open} className="input">
          <div className="swatch">
            <div className="color" style={{ backgroundColor: value }} />
          </div>
          <div className="label">{label}</div>
        </div>
        {displayed ? (
          <div className="popover">
            <div className="cover" onClick={close} />
            <ChromePicker color={value} disableAlpha={true} onChange={onChange} />
          </div>
        ) : null}
      </div>
    )
  }
}

export default ColorPicker
