import React from 'react'
import AppBar from 'react-toolbox/lib/app_bar/AppBar'
import Navigation from 'react-toolbox/lib/navigation/Navigation'
import Icon from 'react-toolbox/lib/font_icon/FontIcon'
import { Link } from "react-router-dom"
import classnames from 'classnames'

import './Header.css'

const menu = [
  { url: '/medias', title: 'Medias', icon: 'photo_library' },
  { url: '/users', title: 'Utilisateurs', icon: 'supervisor_account', role: 'admin' },
  { url: '/plugins', title: 'Plugins', icon: 'widgets', role: 'admin' },
  { url: '/profil', getTitle: user => user.displayName, icon: 'person' },
]

const getMenuProps = ({ isPartiallyCurrent }) => ({
  className: classnames('AppBarLink', { active: isPartiallyCurrent }),
})

const Header = ({ goBack, menuIcon, user, logout }) => {
  let navigation = null
  if (user) {
    const menuItems = menu.filter(({ role }) => !role || user.role === role)
    navigation = (
      <Navigation className="AppBarNavigation">
        {menuItems.map(({ url, getTitle, title, icon }) => (
          <Link className="AppBarLink" to={url} key={url}> {/* FIXME  getProps={getMenuProps} */}
            <Icon>{icon}</Icon> {title ? title : getTitle(user)}
          </Link>
        ))}
      </Navigation>
    )
  }

  return (
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
  )
}

export default Header
