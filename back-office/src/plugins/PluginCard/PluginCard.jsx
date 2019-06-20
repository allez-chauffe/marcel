import React from 'react'
import CardMedia from 'react-toolbox/lib/card/CardMedia'
import CardTitle from 'react-toolbox/lib/card/CardTitle'
import CardText from 'react-toolbox/lib/card/CardText'
import CardActions from 'react-toolbox/lib/card/CardActions'
import Button from 'react-toolbox/lib/button/Button'
import ProgressBar from 'react-toolbox/lib/progress_bar/ProgressBar'
import { Card } from '../../components/commons'

import './PluginCard.css'

const PluginTitle = ({ plugin, updating }) => (
  <>
    <span>{plugin.name}</span>
    <span className="PluginVersion">
      {updating ? (
        <ProgressBar type="circular" mode="indeterminate" className="updating" />
      ) : (
        plugin.version
      )}
    </span>
  </>
)

const UpdateButton = ({ update, updating }) => (
  <Button icon="update" label="mettre à jour" disabled={updating} onClick={update} />
)

const PluginCard = ({ plugin, update, updating }) => (
  <Card clickable={false}>
    <CardMedia aspectRatio="wide" image="https://placeimg.com/800/450/nature" />
    <CardTitle title={<PluginTitle plugin={plugin} updating={updating} />} />
    <CardText>{plugin.description}</CardText>
    <CardActions className="buttons">
      <UpdateButton update={update} updating={updating} />
    </CardActions>
  </Card>
)

export default PluginCard
