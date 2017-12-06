// @flow
import React, { Component } from 'react'
import AppBar from 'react-toolbox/lib/app_bar/AppBar'
import ReduxToastr from 'react-redux-toastr'

import { Auth } from '../../components/auth'
import Dialogs from '../Dialogs'
import HomePage from '../HomePage'
import MediaEditPage from '../MediaEditPage'
import MediaListPage from '../MediaListPage'

import './AppLayout.css'

class AppLayout extends Component {
  props: {
    goBack: () => void,
    menuIcon: string,
  }

  componentDidMount() {}

  render() {
    const { goBack, menuIcon } = this.props
    return (
      <div className="AppLayout">
        <header>
          <AppBar title="Zenboard" leftIcon={menuIcon} onLeftIconClick={goBack} />
        </header>

        <main>
          <HomePage />
          <Auth>
            <MediaListPage />
            <MediaEditPage />
          </Auth>
        </main>

        <Dialogs />
        <ReduxToastr preventDuplicates position="top-right" />
      </div>
    )
  }
}

export default AppLayout
