// @flow
import React from 'react'
import AppBar from 'react-toolbox/lib/app_bar/AppBar'
import MainScreen from '../MainScreen'

import './AppLayout.css'

const AppLayout = () => (
  <div className="AppLayout">
    <header>
      <AppBar title="Zenboard" leftIcon="menu" />
    </header>
    <main>
      <MainScreen />
    </main>
    <footer>
      <AppBar />
    </footer>
  </div>
)

export default AppLayout
