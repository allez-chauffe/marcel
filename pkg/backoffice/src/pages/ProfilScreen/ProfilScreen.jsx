import React from 'react'

import './ProfilScreen.css'
import { NewUserLine } from '../../components/user'

class ProfilScreen extends React.Component {
  render() {
    const { user, handleSave, handleChange } = this.props
    return (
      <div className="ProfilScreen">
        <NewUserLine
          user={user}
          handleSave={handleSave}
          handleChange={handleChange}
          disableInputRole={true}
        />
      </div>
    )
  }
}

export default ProfilScreen
