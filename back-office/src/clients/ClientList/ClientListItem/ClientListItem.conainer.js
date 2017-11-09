//@flow
import { connect } from 'react-redux'
import type { Dispatch } from 'redux'
import ClientListItem from './ClientListItem'
import type { State, Action } from '../../../store'
import { associateClient, requireClientAssociation } from '../../actions'
import { isClientLoadingSelector } from '../../selectors'

const mapStateToProps = (state: State, ownProps: { client: Client }) => ({
  isLoading: isClientLoadingSelector(state, ownProps.client),
})

const mapDispatchToProps = (dispatch: Dispatch<Action>, ownProps) => {
  const { client } = ownProps
  return {
    associate: () =>
      dispatch(client.mediaID ? requireClientAssociation(client) : associateClient(client)),
  }
}

export default connect(mapStateToProps, mapDispatchToProps)(ClientListItem)
