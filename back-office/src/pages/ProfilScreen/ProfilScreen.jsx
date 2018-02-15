//@flow
import React from 'react'

import './ProfilScreen.css'
import { User } from '../../user';
import { NewUserLine } from '../../components/user'

class ProfilScreen extends React.Component {
  
  props: {
    user: User,
    handleSave: () => void,
    handleChange: () => void,
  }

  render() {
    const { user, handleSave, handleChange } = this.props
    return (
      <div className="ProfilScreen">
        <NewUserLine 
          user={user} 
          handleSave={handleSave} 
          handleChange={handleChange} 
          disableInputRole={true} />     
      </div>
    )
  }
}

export default ProfilScreen
