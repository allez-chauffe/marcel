//@flow
import React from 'react'
import { MediaCard, AddMediaCard } from '../../components/media'
import type { Dashboard } from '../../dashboard/type'

import './MediaListPage.css'

class DashboardScreen extends React.Component {
  props: {
    medias: Dashboard[],
  }

  render() {
    return (
      <div className="MediaListPage">
        <AddMediaCard />
        {this.props.medias.map(media => <MediaCard key={media.id} dashboard={media} />)}
      </div>
    )
  }
}

export default DashboardScreen
