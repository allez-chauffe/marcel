//@flow
import React from 'react'
import CardText from 'react-toolbox/lib/card/CardText'
import FontIcon from 'react-toolbox/lib/font_icon/FontIcon'
import DashboardCard from '../DashboardCard'

import './AddDashboardCard.css'

const AddDashboardCard = () =>
  <DashboardCard className="AddDashboardCard">
    <CardText className="addText">
      <FontIcon value="add" className="addIcon" />
    </CardText>
  </DashboardCard>

export default AddDashboardCard
