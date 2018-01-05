import React, { Component } from 'react'
import { toast } from 'react-toastify'
import { backendFetcher } from '../utils/fetcher'
import isEqual from 'lodash/isEqual'
import Plugin from './Plugin'
import Loader from './Loader'

class Media extends Component {
  state = {
    loading: true,
  }

  getMedia = () =>
    this.backend.get(`/medias/${this.props.mediaId}/`).then(media => {
      console.log('Media loaded: ', media)
      return media
    })

  setMedia = media => {
    this.setState({ media, loading: false })
    if (!this.state.media.isactive) this.inactiveMedia()
  }

  componentDidMount() {
    console.log('media mount')
    this.backend = backendFetcher()
    this.getMedia(this.props.mediaId)
      .then(this.setMedia)
      .catch(error => {
        toast.error('Un erreur est survenue lors du chargement du Media', { autoClose: false })
        throw error
      })
  }

  componentDidUpdate(prevProps) {
    if (isEqual(prevProps, this.props)) return
    this.getMedia().then(this.setMedia)
  }

  inactiveMedia() {
    toast.error("Ce Media n'est pas actif", { autoClose: false })
  }

  render() {
    if (this.state.loading) return <Loader />
    if (!this.state.media.isactive) return null

    const { mediaId, config: { ssl, urls } } = this.props

    const col = 100 / this.state.media.cols
    const row = 100 / this.state.media.rows
    const pluginsURL = `http${ssl ? 's' : ''}://${urls.backend}/medias/${mediaId}/plugins`

    return (
      <div className="media fullSize">
        {this.state.media.plugins.map(plugin => (
          <Plugin
            plugin={plugin}
            pluginsURL={pluginsURL}
            key={plugin.instanceId}
            style={{
              width: plugin.frontend.cols * col + '%',
              height: plugin.frontend.rows * row + '%',
              left: plugin.frontend.x * col + '%',
              top: plugin.frontend.y * row + '%',
            }}
          />
        ))}
      </div>
    )
  }
}

export default Media
