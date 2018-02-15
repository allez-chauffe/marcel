//@flow
import React from 'react'
import type { Children } from 'react'
import ReactCard from 'react-toolbox/lib/card/Card'

import './Card.css'

export type PropsType = {
  children: Children,
  style?: { [prop: string]: string },
  className?: string,
  onClick?: MouseEvent => void,
}

const Card = (props: PropsType) => {
  const { children, style, className, onClick } = props
  return (
    <ReactCard onClick={onClick} style={{ ...style }} className={`Card ${className || ''}`}>
      {children}
    </ReactCard>
  )
}

export default Card
