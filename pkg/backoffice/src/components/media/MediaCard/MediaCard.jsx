import React, { Component } from 'react'
import CardMedia from 'react-toolbox/lib/card/CardMedia'
import CardTitle from 'react-toolbox/lib/card/CardTitle'
import CardText from 'react-toolbox/lib/card/CardText'
import CardActions from 'react-toolbox/lib/card/CardActions'
import Button from 'react-toolbox/lib/button/Button'
import { ActivationButton, OpenButton, DeleteDashboardButton } from '../../../common'
import { Card } from '../../commons'

import './MediaCard.css'

class DashboardListItem extends Component {
  onClickWithoutPropagation = onClick => event => {
    event.stopPropagation()
    onClick()
  }

  selectDashboard = this.onClickWithoutPropagation(this.props.selectDashboard)
  deleteDashboard = this.onClickWithoutPropagation(this.props.deleteDashboard)
  openDashboard = this.onClickWithoutPropagation(() => {
    const { Frontend, dashboard } = this.props
    window.open(Frontend + dashboard.id)
    window.focus()
  })

  render() {
    const { dashboard } = this.props
    const { selectDashboard } = this
    return (
      <Card onClick={selectDashboard}>
        <CardMedia aspectRatio="wide" image="https://placeimg.com/800/450/nature" />
        <CardTitle title={dashboard.name} />
        <CardText>{dashboard.description}</CardText>
        <CardActions className="buttons">
          <Button icon="mode_edit" label="modifier" onClick={selectDashboard} />
          <OpenButton dashboard={dashboard} />
          <DeleteDashboardButton dashboard={dashboard} />
          <ActivationButton dashboard={dashboard} />
        </CardActions>
      </Card>
    )
  }
}

export default DashboardListItem
