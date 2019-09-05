import React from 'react'
import { Route, Redirect, BrowserRouter as Router } from "react-router-dom"
import ReduxToastr from 'react-redux-toastr'

import Header from './Header'
import { Auth } from '../../components/auth'
import Dialogs from '../Dialogs'
import MediaEditPage from '../MediaEditPage'
import MediaListPage from '../MediaListPage'
import UserScreen from '../UserScreen'
import ProfilScreen from '../ProfilScreen'
import PluginsScreen from '../PluginsScreen'

import './AppLayout.css'


const AppLayout = () => (
  <div className="AppLayout">
    <Router>
      <Header />

      <main>
        <Redirect from="/" to="/medias" />
        <Auth>
          <Route path="/medias" exact component={MediaListPage} />
          <Route path="/medias/:mediaId" component={MediaEditPage} />
          <Route path="/users" component={UserScreen} />
          <Route path="/profil" component={ProfilScreen} />
          <Route path="/plugins" component={PluginsScreen} />
        </Auth>
      </main>
    </Router>

    <Dialogs />
    <ReduxToastr preventDuplicates position="top-right" />
  </div>
)

export default AppLayout
