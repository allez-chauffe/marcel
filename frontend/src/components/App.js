import React, { Component } from 'react'
import { ToastContainer, toast } from 'react-toastify'
import { localFetcher } from '../utils/fetcher'
import Client from './Client'
import Loader from './Loader'
import Auth from './Auth'

class App extends Component {
  state = {
    loading: true,
  }

  getConfig = async () => {
    try {
      const config = await localFetcher.get('./config')
      console.log('Local config loaded', config)
      return config
    } catch (e) {
      if (e.status === 404) {
        console.warn('No config found')
        return { API: '/api/' }
      }
      throw e
    }
  }

  componentDidMount() {
    this.getConfig()
      .then(config => this.setState({ config, loading: false }))
      .catch(error => {
        toast.error('Erreur lors du chargement de la configuration', { autoClose: false })
        throw error
      })
  }

  render() {
    return (
      <div className="fullSize">
        {this.state.loading && <Loader />}
        {this.state.loading || (
          <Auth config={this.state.config}>
            <Client />
          </Auth>
        )}
        <ToastContainer closeButton={false} hideProgressBar closeOnClick />
      </div>
    )
  }
}

export default App
