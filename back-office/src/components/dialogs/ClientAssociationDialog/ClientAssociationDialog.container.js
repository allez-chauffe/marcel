import { connect } from 'react-redux'
import {
  associatingClientSelector,
  confirmClientAssociation,
  cancelClientAssociation,
  associateClient,
} from '../../../clients'
import { dashboardsSelector } from '../../../dashboard'
import AssociationClientDialog from './ClientAssociationDialog'

const mapStateToProps = state => {
  const associating = associatingClientSelector(state)
  return {
    associating,
    media: associating && dashboardsSelector(state)[associating.mediaID],
  }
}

const mapDispatchToProps = {
  confirm: confirmClientAssociation,
  cancel: cancelClientAssociation,
  associate: associateClient,
}

export default connect(
  mapStateToProps,
  mapDispatchToProps,
)(AssociationClientDialog)
