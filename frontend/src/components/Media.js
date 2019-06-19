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
    const { media, loading } = this.state

    if (loading) return <Loader />
    if (!media.isactive) return null

    const {
      mediaId,
      config: { urls },
    } = this.props

    const col = 100 / media.cols
    const row = 100 / media.rows
    const pluginsURL = `${urls.plugins}/medias/${mediaId}/plugins`

    const styles = {
      backgroundColor: media.stylesvar['background-color'],
      color: media.stylesvar['primary-color'],
      fontFamily: media.stylesvar['font-family'],
    }

    return (
      <div className="media fullSize" style={styles}>
        {media.plugins.map(plugin => (
          <Plugin
            plugin={plugin}
            stylesvar={media.stylesvar}
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
