//@flow
import React from 'react'
import Dialog from 'react-toolbox/lib/dialog/Dialog'

export type PropsType = {
  isDeletingDashboard: boolean,
  confirmDeletion: () => void,
  cancelDeletion: () => void,
}

const DeleteDahboardDialog = (props: PropsType) => {
  const { isDeletingDashboard, confirmDeletion, cancelDeletion } = props

  return (
    <Dialog
      title="Etes-vous sûre de vouloir supprimer ce dahsboard ?"
      type="small"
      active={isDeletingDashboard}
      onEscKeyDown={cancelDeletion}
      onOverlayClick={cancelDeletion}
      actions={[
        { label: 'Anuler', onClick: cancelDeletion },
        { label: 'Supprimer', onClick: confirmDeletion },
      ]}
    >
      Attention, Cette action est définitive et ne pourra pas être annulée.
    </Dialog>
  )
}

export default DeleteDahboardDialog
