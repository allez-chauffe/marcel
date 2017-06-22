//@flow
import React from 'react'
import Card from 'react-toolbox/lib/card/Card'
import CardMedia from 'react-toolbox/lib/card/CardMedia'
import CardTitle from 'react-toolbox/lib/card/CardTitle'
import CardText from 'react-toolbox/lib/card/CardText'
import CardActions from 'react-toolbox/lib/card/CardActions'
import Button from 'react-toolbox/lib/button/Button'
import type { Dashboard } from '../../../../dashboard/type'

import './DashboardListItem.css'

export type PropsType = {
  dashboard: Dashboard,
  selectDashboard: () => void,
}

const DashboardListItem = (props: PropsType) => {
  const { dashboard, selectDashboard } = props
  return (
    <Card style={{ width: '400px' }} className="DashboardListItem">
      <CardMedia
        aspectRatio="wide"
        image="https://placeimg.com/800/450/nature"
      />
      <CardTitle title={dashboard.name} />
      <CardText>{dashboard.description}</CardText>
      <CardActions>
        <Button label="modifier" onClick={selectDashboard} />
        <Button label="ouvrir" />
        <Button label="supprimer" />
      </CardActions>
    </Card>
  )
}

export default DashboardListItem
