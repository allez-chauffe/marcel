import React from 'react'
import { MediaCard, AddMediaCard } from '../../components/media'

import './MediaListPage.css'

class DashboardScreen extends React.Component {
  render() {
    return (
      <div className="MediaListPage CardGrid">
        <AddMediaCard />
        {this.props.medias.map(media => (
          <MediaCard key={media.id} dashboard={media} />
        ))}
      </div>
    )
  }
}

export default DashboardScreen
