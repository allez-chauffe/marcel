import React from 'react'
import { bool } from 'prop-types'

import './OccupationIndicator.css'

const OccupationIndicator = props => {
  const { isOccupied } = props
  return (
    <div className="OccupationIndicator">
      <div className={`indicator ${isOccupied ? 'occupied' : ''}`} />
    </div>
  )
}

OccupationIndicator.propTypes = {
  isOccupied: bool,
}

export default OccupationIndicator
