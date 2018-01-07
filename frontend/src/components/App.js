import React, { Component } from 'react'
import { ToastContainer, toast } from 'react-toastify'
import { localFetcher } from '../utils/fetcher'
import Client from './Client'
import Loader from './Loader'

class App extends Component {
  state = {
    loading: true,
  }

  getConfig = () =>
    localFetcher.get('conf/config.json').then(config => {
      console.log('Local config loaded', config)
      return config
    })

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
        {this.state.loading || <Client config={this.state.config} />}
        <ToastContainer closeButton={false} hideProgressBar closeOnClick />
      </div>
    )
  }
}

export default App
