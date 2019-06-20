import React from 'react'

import './UserScreen.css'
import { UserLine, NewUserLine } from '../../components/user'

class UserScreen extends React.Component {
  render() {
    const {
      users,
      userEdited,
      editUserHandleChange,
      editUserHandleSave,
      editUserHandleReset,
    } = this.props
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
          <tbody>{users ? users.map(user => <UserLine user={user} key={user.id} />) : ''}</tbody>
        </table>

        <NewUserLine
          className="UserForm"
          user={userEdited}
          handleSave={editUserHandleSave}
          handleChange={editUserHandleChange}
          handleReset={editUserHandleReset}
        />
      </div>
    )
  }
}

export default UserScreen
