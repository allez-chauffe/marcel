//@flow
import React from 'react'

import './UserScreen.css'
import { User } from '../../user';
import { UserLine, NewUserLine } from '../../components/user'

class UserScreen extends React.Component {
  
  props: {
    users: User[],
    userEdited: User,
    addUser: any,
    updateUser: any,
    updateCurrentUserProperty: any,
    resetCurrentUser: any,
  }

  componentWillMount() {
  }

  editUserHandleChange = (name, value) => {
    this.props.updateCurrentUserProperty(name, value)
  }

  editUserHandleSave = (user) => {
    if (user.id) {
      this.props.updateUser(user)
    } else {
      this.props.addUser(user)
    } 
  }

  editUserHandleReset = () => {
    this.props.resetCurrentUser()
  }

  render() {
    const { users, userEdited } = this.props
    return (
      <div className="UserScreen">

        <table className="UsersList">
          <thead>
            <tr>
              <td>Nom</td>
              <td>Login</td>
              <td>Rôle</td>
              <td>Action</td>
            </tr>
          </thead>
          <tbody>
            {users ? users.map(user => <UserLine user={user} key={user.id}/>) : ''}
          </tbody>
        </table>

        <NewUserLine className="UserForm" user={userEdited} handleSave={this.editUserHandleSave} handleChange={this.editUserHandleChange} handleReset={this.editUserHandleReset} />
      </div>
    )
  }
}

export default UserScreen
