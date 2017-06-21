// @flow
import React from 'react'
import AppBar from 'react-toolbox/lib/app_bar/AppBar'
import DashboardScreen from '../DashboardScreen'

import './AppLayout.css'

const AppLayout = () =>
  <div className="AppLayout">
    <header>
      <AppBar title="Zenboard" leftIcon="menu" />
    </header>
    <main>
      <DashboardScreen />
    </main>
    <footer>
      <AppBar />
    </footer>
  </div>

export default AppLayout
