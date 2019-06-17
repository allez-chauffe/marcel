import React from 'react'
import { shape, string, bool, func } from 'prop-types'
import ListItem from 'react-toolbox/lib/list/ListItem'
import IconButton from 'react-toolbox/lib/button/IconButton'
import ProgressBar from 'react-toolbox/lib/progress_bar/ProgressBar'
import OccupationIndicator from './OccupationIndicator'

import './ClientListItem.css'

const ClientListItem = (props) => {
  const { client, associate, isLoading } = props
  const { name, id, type, mediaID } = client
  const mediaName = mediaID ? ` (${mediaID})` : ''
  return (
    <div className="ClientListItem">
      <ListItem
        caption={name + mediaName}
        ripple={false}
        key={id}
        legend={type}
        leftIcon={<OccupationIndicator isOccupied={!!mediaID} />}
        rightIcon={
          isLoading ? (
            <ProgressBar type="circular" />
          ) : (
            <IconButton icon="open_in_browser" onClick={associate} />
          )
        }
      />
    </div>
  )
}

ClientListItem.propTypes = {
  client: shape({
    name: string,
    id: string,
    type: string,
    mediaId: string,
  }),
  associate: func,
  isLoading: bool,
}

export default ClientListItem
