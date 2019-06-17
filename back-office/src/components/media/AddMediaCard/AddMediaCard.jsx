import React from 'react'
import CardText from 'react-toolbox/lib/card/CardText'
import FontIcon from 'react-toolbox/lib/font_icon/FontIcon'
import { Card } from '../../commons'

import './AddMediaCard.css'

const AddMediaCard = props => {
  const { addDashboard } = props
  return (
    <div onClick={addDashboard}>
      <Card className="AddMediaCard">
        <CardText className="addText">
          <FontIcon value="add" className="addIcon" />
        </CardText>
      </Card>
    </div>
  )
}

export default AddMediaCard
