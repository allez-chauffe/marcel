import React, { Component } from 'react'
import Dialog from 'react-toolbox/lib/dialog/Dialog'

class ClientAssociationDialog extends Component {
  confirm = () => {
    this.props.confirm()
    this.props.associating && this.props.associate(this.props.associating)
  }

  render() {
    const { associating, cancel, media } = this.props
    const clientName = associating ? associating.name : ''
    const mediaName = media ? media.name : ''

    return (
      <Dialog
        title={`Le client ${clientName} affiche déjà le média ${mediaName}. Êtes-vous sûre de vouloir changer ?`}
        type="small"
        active={!!associating}
        onEscKeyDown={cancel}
        onOverlayClick={cancel}
        actions={[
          { label: 'Anuler', onClick: cancel, icon: 'cancel' },
          {
            label: 'Associer',
            onClick: this.confirm,
            icon: 'open_in_browser',
          },
        ]}
      >
        Attention, Cette action est définitive et ne pourra pas être annulée.
      </Dialog>
    )
  }
}

export default ClientAssociationDialog
