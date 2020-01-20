import React, { Component } from 'react'
import { ToastContainer, toast } from 'react-toastify'
import { rootFetcher } from '../utils/fetcher'
import Client from './Client'
import Loader from './Loader'
import Auth from './Auth'

class App extends Component {
  state = {
    loading: true,
  }

  getURIs = async () => {
    try {
      return await rootFetcher.get('/uris')
    } catch (e) {
      if (e.status === 404) {
        console.warn('No URIs config found')
        return { API: '/api/' }
      }
      throw e
    }
  }

  componentDidMount() {
    this.getURIs()
      .then(uris => this.setState({ uris, loading: false }))
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
          <Auth uris={this.state.uris}>
            <Client />
          </Auth>
        )}
        <ToastContainer closeButton={false} hideProgressBar closeOnClick />
      </div>
    )
  }
}

export default App
