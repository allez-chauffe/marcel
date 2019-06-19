import React from 'react'
import AppBar from 'react-toolbox/lib/app_bar/AppBar'
import Navigation from 'react-toolbox/lib/navigation/Navigation'
import Icon from 'react-toolbox/lib/font_icon/FontIcon'
import { Redirect, Link } from '@reach/router'
import classnames from 'classnames'

import ReduxToastr from 'react-redux-toastr'

import { Auth } from '../../components/auth'
import Dialogs from '../Dialogs'
import MediaEditPage from '../MediaEditPage'
import MediaListPage from '../MediaListPage'
import UserScreen from '../UserScreen'
import ProfilScreen from '../ProfilScreen'

import './AppLayout.css'
import { Router } from '@reach/router'
import PluginsScreen from '../PluginsScreen'

const menu = [
  { url: '/medias', title: 'Medias', icon: 'photo_library' },
  { url: '/users', title: 'Utilisateurs', icon: 'supervisor_account', role: 'admin' },
  { url: '/plugins', title: 'Plugins', icon: 'widgets', role: 'admin' },
  { url: '/profil', getTitle: user => user.displayName, icon: 'person' },
]

const getMenuProps = ({ isPartiallyCurrent }) => ({
  className: classnames('AppBarLink', { active: isPartiallyCurrent }),
})

const AppLayout = ({ goBack, menuIcon, user, logout }) => {
  let navigation = null
  if (user) {
    const menuItems = menu.filter(({ role }) => !role || user.role === role)
    navigation = (
      <Navigation className="AppBarNavigation">
        {menuItems.map(({ url, getTitle, title, icon }) => (
          <Link className="AppBarLink" to={url} key={url} getProps={getMenuProps}>
            <Icon>{icon}</Icon> {title ? title : getTitle(user)}
          </Link>
        ))}
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
            <PluginsScreen path="/plugins" />
          </Auth>
        </Router>
      </main>

      <Dialogs />
      <ReduxToastr preventDuplicates position="top-right" />
    </div>
  )
}

export default AppLayout
