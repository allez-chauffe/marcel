//@flow
import React from 'react'
import CardText from 'react-toolbox/lib/card/CardText'
import FontIcon from 'react-toolbox/lib/font_icon/FontIcon'
import DashboardCard from '../DashboardCard'

import './AddDashboardCard.css'

export type PropsType = {
  addDashboard: () => void,
}

const AddDashboardCard = (props: PropsType) => {
  const { addDashboard } = props
  return (
    <div onClick={addDashboard}>
      <DashboardCard className="AddDashboardCard">
        <CardText className="addText">
          <FontIcon value="add" className="addIcon" />
        </CardText>
      </DashboardCard>
    </div>
  )
}

export default AddDashboardCard
