import React, { Component } from 'react'
import Button from 'react-toolbox/lib/button/Button'

class DeleteDashboardButton extends Component {
  delete = event => {
    event.stopPropagation()
    this.props.delete()
  }

  render = () => {
    if (!this.props.dashboard.isWritable) return null
    return <Button icon="delete" label="supprimer" onClick={this.delete} />
  }
}

export default DeleteDashboardButton
