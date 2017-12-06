//@flow
import React from 'react'
import Dialog from 'react-toolbox/lib/dialog/Dialog'

export type PropsType = {
  isDeletingDashboard: boolean,
  confirmDeletion: () => void,
  cancelDeletion: () => void,
}

const MediaDeletionDialog = (props: PropsType) => {
  const { isDeletingDashboard, confirmDeletion, cancelDeletion } = props

  return (
    <Dialog
      title="Etes-vous sûre de vouloir supprimer ce Media ?"
      type="small"
      active={isDeletingDashboard}
      onEscKeyDown={cancelDeletion}
      onOverlayClick={cancelDeletion}
      actions={[
        { label: 'Anuler', onClick: cancelDeletion, icon: 'cancel' },
        {
          label: 'Supprimer',
          onClick: confirmDeletion,
          icon: 'delete_forever',
        },
      ]}
    >
      Attention, Cette action est définitive et ne pourra pas être annulée.
    </Dialog>
  )
}

export default MediaDeletionDialog
