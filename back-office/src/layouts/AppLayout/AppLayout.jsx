// @flow
import React from 'react'
import AppBar from 'react-toolbox/lib/app_bar/AppBar'
import DashboardScreen from '../DashboardScreen'
import { Auth } from '../../auth'
import AssociationClientDialog from './AssociationClientDialog'

import './AppLayout.css'

export type PropsType = {
  unselectDashboard: () => void,
  isDashboardSelected: boolean,
}

const AppLayout = (props: PropsType) => {
  const { unselectDashboard, isDashboardSelected } = props
  return (
    <div className="AppLayout">
      <header>
        <AppBar
          title="Zenboard"
          leftIcon={isDashboardSelected ? 'arrow_back' : null}
          onLeftIconClick={unselectDashboard}
        />
      </header>
      <main>
        <Auth>
          <DashboardScreen />
        </Auth>
      </main>
      <AssociationClientDialog />
    </div>
  )
}

export default AppLayout
