//@flow
import React from 'react'
import { ChromePicker } from 'react-color'

import './ColorPicker.css'

class ColorPicker extends React.Component {
  props: {
    label: string,
    value: string,
    onChange: string => void,
  }

  state = {
    displayed: false,
    color: '#700',
    oldColor: '',
  }

  onKeypressed = (event: KeyboardEvent) => {
    if (event.keyCode === 13) this.close() // Enter pressed
    if (event.keyCode === 27) {
      this.onChange({ hex: this.state.oldColor })
      this.close()
    }
  }

  open = () => {
    document.addEventListener('keyup', this.onKeypressed)
    this.setState({ oldColor: this.state.color, displayed: true })
  }

  close = () => {
    document.removeEventListener('keyup', this.onKeypressed)
    this.setState({ displayed: false })
  }

  onChange = (newColor: { hex: string }) =>
    this.setState({ color: newColor.hex })

  render() {
    const { label } = this.props
    const { color, displayed } = this.state
    const { onChange, close, open } = this
    return (
      <div className="ColorPicker">
        <div onClick={open} className="input">
          <div className="swatch">
            <div className="color" style={{ backgroundColor: color }} />
          </div>
          <div className="label">
            {label}
          </div>
        </div>
        {displayed
          ? <div className="popover">
              <div className="cover" onClick={close} />
              <ChromePicker
                color={color}
                disableAlpha={true}
                onChange={onChange}
              />
            </div>
          : null}
      </div>
    )
  }
}

export default ColorPicker
