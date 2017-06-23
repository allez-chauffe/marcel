//@flow
import React from 'react'
import type { Children } from 'react'
import Card from 'react-toolbox/lib/card/Card'

import './DashboardCard.css'

export type PropsType = {
  children: Children,
  style?: { [prop: string]: string },
  className?: string,
}

const DashboardCard = (props: PropsType) => {
  const { children, style, className } = props
  return (
    <Card
      style={{ width: '400px', ...style }}
      className={`DashboardCard ${className || ''}`}
    >
      {children}
    </Card>
  )
}

export default DashboardCard
