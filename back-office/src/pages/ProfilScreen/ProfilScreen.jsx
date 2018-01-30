//@flow
import React from 'react'

import './ProfilScreen.css'
import { User } from '../../user';
import { NewUserLine } from '../../components/user'

class ProfilScreen extends React.Component {
  
  props: {
    user: User,
    updateUser: () => void,
    updateUserProperty: () => void,
  }

  handleSave = (user) => {
    this.props.updateUser(user)
  }

  handleChange = (name, value) => {
    this.props.updateUserProperty(name, value)
  }

  render() {
    const { user } = this.props
    return (
      <div className="ProfilScreen">
        <NewUserLine 
          user={user} 
          handleSave={this.handleSave} 
          handleChange={this.handleChange} 
          disableInputRole={true} />     
      </div>
    )
  }
}

export default ProfilScreen
