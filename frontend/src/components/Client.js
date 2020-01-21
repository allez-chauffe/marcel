import React, { Component } from 'react'
import { toast } from 'react-toastify'
import { backendFetcher } from '../utils/fetcher'
import Media from './Media'
import Loader from './Loader'

class Client extends Component {
  state = {
    loading: true,
  }

  getClient = () =>
    this.backend
      .get(`clients/${this.state.client.id}/`)
      //Checking if a media ID has specified in query params. If any, overrides the server configuration
      .then(client => {
        const name = window.location.queryParams.name
        const mediaID = parseInt(window.location.queryParams.mediaID, 10)

        let newClient = client

        if (name && name !== client.name) newClient = { ...newClient, name }
        if (mediaID && mediaID !== client.mediaID) newClient = { ...newClient, mediaID }

        return newClient === client ? client : this.backend.put(`clients/`, newClient)
      })
      .then(client => {
        console.log('Client loaded', client)
        return client
      })

  getClientId = () => {
    const { clientID } = localStorage

    if (clientID) {
      console.log('Client id found in local storage :', clientID)
      return Promise.resolve(clientID)
    }

    console.log('Creating client id...')

    return this.backend
      .post('clients/', {
        name: window.location.queryParams.name,
        mediaID: parseInt(window.location.queryParams.mediaID, 10)
      })
      .then(client => {
        localStorage.clientID = client.id
        console.log('New client id created by server : ', client.id)
        return client.id
      })
  }

  openWebsocket = () => {
    this.conn = this.backend.ws(this.state.client.id)
    this.conn.onmessage = event => {
      console.log('Message received from server : ', event)
      if (event.data === 'update') {
        toast.dismiss()
        toast.warn('Ce Media a été mis à jour')
        this.getClient().then(this.setClient)
      }
    }

    this.conn.onopen = () => {
      console.log('Connection with server established')
      if (this.state.client.mediaID) toast.dismiss()
    }

    this.conn.onclose = () => {
      setTimeout(() => this.openWebsocket(), 5000)
      toast.error('Connection au backend interrompue .\nTentative de reconnexion dans 5s')
    }
  }

  setClient = client => {
    console.log()
    this.setState({ client, loading: false, lastUpdate: Date.now() })
    if (!client.mediaID) this.noMediaAssociated()
  }

  getNewClientId = () => {
    localStorage.clear()
    return this.getClientId()
  }

  noMediaAssociated = () => {
    toast.info("Aucun Media n'est associé pour ce client", {
      autoClose: false,
      closeOnClick: false,
    })
  }

  componentDidMount() {
    this.backend = backendFetcher(this.props.uris)
    this.getClientId()
      .then(clientId => this.setState({ client: { id: clientId } }))
      .then(this.getClient)
      .catch(error => {
        //If the client doesn't exists anymore, recreate one
        if (error.status === 404) this.getNewClientId().then(this.getClient)
        else throw error
      })
      .then(this.setClient)
      .catch(error => {
        toast.error('Erreur lors du chargement du client', { autoClose: false })
        throw error
      })
      .then(this.openWebsocket)
  }

  componentWillUnmount() {
    this.conn.close()
  }

  render() {
    const { loading, client, lastUpdate } = this.state
    const { uris } = this.props

    if (loading) return <Loader />

    if (client && client.mediaID)
      return <Media uris={uris} mediaId={client.mediaID} lastUpdate={lastUpdate} />

    return null
  }
}

export default Client
