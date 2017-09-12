//@flow
import React from 'react'
import type { Children } from 'react'
import Card from 'react-toolbox/lib/card/Card'

import './DashboardCard.css'

export type PropsType = {
  children: Children,
  style?: { [prop: string]: string },
  className?: string,
  onClick?: MouseEvent => void,
}

const DashboardCard = (props: PropsType) => {
  const { children, style, className, onClick } = props
  return (
    <Card
      onClick={onClick}
      style={{ ...style }}
      className={`DashboardCard ${className || ''}`}
    >
      {children}
    </Card>
  )
}

export default DashboardCard
