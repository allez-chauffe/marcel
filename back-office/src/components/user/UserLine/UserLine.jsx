//@flow
import React from 'react'
import { User } from '../../../user/index'
import IconButton from 'react-toolbox/lib/button/IconButton'

import './UserLine.css'

class UserLine extends React.Component {
  
  props: {
    user: User,
    handleEdit: any,
    handleDelete: any,
  }

  render() {
    const { user, handleEdit, handleDelete } = this.props
    return (
      <tr>
        <td>{user.displayName}</td>
        <td>{user.login}</td>
        <td>{user.role}</td>
        <td className="ActionsCol"><IconButton icon='mode_edit' onClick={handleEdit}/><IconButton icon='delete' onClick={handleDelete} /></td>
      </tr>
    )
  }
}

export default UserLine
