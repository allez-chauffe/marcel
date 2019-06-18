import { connect } from 'react-redux'
import ClientListItem from './ClientListItem'
import { associateClient, requireClientAssociation } from '../../actions'
import { isClientLoadingSelector } from '../../selectors'

const mapStateToProps = (state, ownProps) => ({
  isLoading: isClientLoadingSelector(state, ownProps.client),
})

const mapDispatchToProps = (dispatch, ownProps) => {
  const { client } = ownProps
  return {
    associate: () =>
      dispatch(client.mediaID ? requireClientAssociation(client) : associateClient(client)),
  }
}

export default connect(
  mapStateToProps,
  mapDispatchToProps,
)(ClientListItem)
