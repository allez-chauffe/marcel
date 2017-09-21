//@flow
import { connect } from 'react-redux'
import type { Dispatch } from 'redux'
import ClientListItem from './ClientListItem'
import type { State, Action } from '../../../store'
import { associateClient } from '../../actions'
import { isClientLoadingSelector } from '../../selectors'

const mapStateToProps = (state: State, ownProps: { client: Client }) => ({
  isLoading: isClientLoadingSelector(state, ownProps.client),
})

const mapDispatchToProps = (dispatch: Dispatch<Action>, ownProps) => ({
  associate: () => dispatch(associateClient(ownProps.client)),
})

export default connect(mapStateToProps, mapDispatchToProps)(ClientListItem)
