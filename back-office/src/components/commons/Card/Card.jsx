import React from 'react'
import ReactCard from 'react-toolbox/lib/card/Card'

import './Card.css'

const Card = props => {
  const { children, style, className, onClick } = props
  return (
    <ReactCard onClick={onClick} style={{ ...style }} className={`Card ${className || ''}`}>
      {children}
    </ReactCard>
  )
}

export default Card
