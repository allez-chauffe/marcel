//@flow
import React from 'react'
import Button from 'react-toolbox/lib/button/Button'

class DeleteDashboardButton extends React.Component {
  delete = (event: Event) => {
    event.stopPropagation()
    this.props.delete()
  }

  render = () => {
    if (!this.props.dashboard.isWritable) return null
    return <Button icon="delete" label="supprimer" onClick={this.delete} />
  }
}

export default DeleteDashboardButton
