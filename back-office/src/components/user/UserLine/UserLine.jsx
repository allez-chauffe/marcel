//@flow
import React from 'react'
import { User } from '../../../user/index'
import IconButton from 'react-toolbox/lib/button/IconButton'

import './UserLine.css'

class UserLine extends React.Component {
  
  props: {
    user: User,
    editUser: any,
    deleteUser: any,
  }

  handleEdit = () => {
    this.props.editUser(this.props.user)
  }

  handleDelete = () => {
    this.props.deleteUser(this.props.user.id)
  }

  render() {
    const { user } = this.props
    return (
      <tr>
        <td>{user.displayName}</td>
        <td>{user.login}</td>
        <td>{user.role}</td>
        <td className="ActionsCol"><IconButton icon='mode_edit' onClick={this.handleEdit}/><IconButton icon='delete' onClick={this.handleDelete} /></td>
      </tr>
    )
  }
}

export default UserLine
