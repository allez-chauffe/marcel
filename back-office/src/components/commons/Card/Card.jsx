import React from 'react'
import ReactCard from 'react-toolbox/lib/card/Card'
import classnames from 'classnames'

import './Card.css'

const Card = props => {
  const { children, style, className, onClick, clickable = true } = props
  return (
    <ReactCard
      onClick={onClick}
      style={{ ...style }}
      className={classnames('Card', className, { clickable })}
    >
      {children}
    </ReactCard>
  )
}

export default Card
