//@flow
import React from 'react'
import List from 'react-toolbox/lib/list/List'
import ListSubHeader from 'react-toolbox/lib/list/ListSubHeader'
import ClientListItem from './ClientListItem'
import { SearchField } from '../../common'
import type { Client } from '../type'

import './ClientList.css'

type PropTypes = {
  clients: {
    associated: Client[],
    connected: Client[],
    disconnected: Client[],
  },
  filter: string,
  changeFilter: string => void,
}

const ClientList = (props: PropTypes) => {
  const { clients, filter, changeFilter } = props
  const { associated, connected, disconnected } = clients

  return (
    <div className="ClientList">
      <SearchField label="Rechercher un client" value={filter} onChange={changeFilter} />
      <List selectable>
        <ListSubHeader caption="Clients associés" />
        {associated ? (
          associated.map(c => <ClientListItem client={c} key={c.id} />)
        ) : (
          <div className="emptyLabel">Aucun client associé</div>
        )}

        <ListSubHeader caption="Clients connectés" />
        {connected ? (
          connected.map(c => <ClientListItem client={c} key={c.id} />)
        ) : (
          <div className="emptyLabel">Aucun client connecté</div>
        )}

        <ListSubHeader caption="Client déconnectés" />
        {disconnected ? (
          disconnected.map(c => <ClientListItem client={c} key={c.id} />)
        ) : (
          <div className="emptyLabel">Aucun client déconnecté</div>
        )}
      </List>
    </div>
  )
}

export default ClientList
