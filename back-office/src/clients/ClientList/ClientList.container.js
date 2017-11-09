//@flow
import { connect } from 'react-redux'
import ClientList from './ClientList'
import {
  clientsFilterSelector,
  partionedFilteredClientsSelector,
  changeClientsFilter,
} from '../../store/filters'
import type { State } from '../../store'

const mapStateToProps = (state: State) => ({
  clients: partionedFilteredClientsSelector(state),
  filter: clientsFilterSelector(state),
})

const mapDispatchToProps = {
  changeFilter: changeClientsFilter,
}

export default connect(mapStateToProps, mapDispatchToProps)(ClientList)
