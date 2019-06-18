import React from 'react'
import AppBar from 'react-toolbox/lib/app_bar/AppBar'
import Navigation from 'react-toolbox/lib/navigation/Navigation'
import Icon from 'react-toolbox/lib/font_icon/FontIcon'
import { Redirect, Link } from '@reach/router'

import ReduxToastr from 'react-redux-toastr'

import { Auth } from '../../components/auth'
import Dialogs from '../Dialogs'
import MediaEditPage from '../MediaEditPage'
import MediaListPage from '../MediaListPage'
import UserScreen from '../UserScreen'
import ProfilScreen from '../ProfilScreen'

import './AppLayout.css'
import { Router } from '@reach/router'

const AppLayout = ({ goBack, menuIcon, user, logout }) => {
  let navigation = ''
  if (user) {
    navigation = (
      <Navigation className="AppBarNavigation">
        <Link className="AppBarLink" to="/medias">
          <Icon>photo_library</Icon>
          Medias
        </Link>
        {user.role === 'admin' && (
          <Link className="AppBarLink" to="/users">
            <Icon>supervisor_account</Icon>
            Utilisateurs
          </Link>
        )}
        <Link className="AppBarLink" to="/profil">
          <Icon>person</Icon>
          {user.displayName}
        </Link>
      </Navigation>
    )
  }

  return (
    <div className="AppLayout">
      <header>
        <AppBar
          title="Zenboard"
          leftIcon={menuIcon}
          onLeftIconClick={goBack}
          rightIcon={user && 'power_settings_new'}
          onRightIconClick={logout}
        >
          {navigation}
        </AppBar>
      </header>

      <main>
        <Router>
          <Redirect from="/" to="/medias" noThrow />
          <Auth path="/">
            <MediaListPage path="/medias" />
            <MediaEditPage path="/medias/:mediaId" />
            <UserScreen path="/users" />
            <ProfilScreen path="/profil" />
          </Auth>
        </Router>
      </main>

      <Dialogs />
      <ReduxToastr preventDuplicates position="top-right" />
    </div>
  )
}

export default AppLayout
