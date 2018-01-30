// @flow
import React, { Component } from 'react'
import AppBar from 'react-toolbox/lib/app_bar/AppBar'
import Navigation from 'react-toolbox/lib/navigation/Navigation'
import Link from 'react-toolbox/lib/link/Link'

import ReduxToastr from 'react-redux-toastr'

import { Auth } from '../../components/auth'
import Dialogs from '../Dialogs'
import HomePage from '../HomePage'
import MediaEditPage from '../MediaEditPage'
import MediaListPage from '../MediaListPage'
import UserScreen from '../UserScreen'
import ProfilScreen from '../ProfilScreen'
import { User } from '../../user'


import './AppLayout.css'

class AppLayout extends Component {
  props: {
    goBack: () => void,
    menuIcon: string,
    user: User,
    logout: () => void,
  }

  componentDidMount() {}

  render() {
    const { goBack, menuIcon, user, logout } = this.props
    const currentPath = window.location.pathname.slice(1)

    let adminLink = ''
    let navigation = ''
    let lougoutIcon = ''
    if (user) {
      lougoutIcon = 'power_settings_new'
      if (user.role === 'admin') {
        adminLink = <Link className="AppBarLink" href='/users'  active={window.location.pathname.slice(1) === 'users'} label='Utilisateurs' icon='supervisor_account' />
      }

      navigation = 
      <Navigation className="AppBarNavigation">
        <Link className="AppBarLink" href='/medias' active={currentPath === 'medias' } label='Medias' icon='photo_library' />
        {adminLink}
        <Link className="AppBarLink" href='/profil' active={currentPath === 'profil'} label={user.displayName} icon='person' />
      </Navigation>
    }
    

    return (
      <div className="AppLayout">
        <header>
          <AppBar title="Zenboard" leftIcon={menuIcon} onLeftIconClick={goBack} rightIcon={lougoutIcon} onRightIconClick={logout} >
            {navigation}
          </AppBar>
        </header>

        <main>
          <HomePage />
          <Auth>
            <MediaListPage />
            <MediaEditPage />
            <UserScreen />
            <ProfilScreen />
          </Auth>
        </main>

        <Dialogs />
        <ReduxToastr preventDuplicates position="top-right" />
      </div>
    )
  }
}

export default AppLayout
