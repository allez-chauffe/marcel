//@flow
import React from 'react'
import Button from 'react-toolbox/lib/button/Button'

class DeleteDashboardButton extends React.Component {
  props: {
    delete: () => void,
  }

  delete = (event: Event) => {
    event.stopPropagation()
    this.props.delete()
  }

  render = () => <Button icon="delete" label="supprimer" onClick={this.delete} />
}

export default DeleteDashboardButton
