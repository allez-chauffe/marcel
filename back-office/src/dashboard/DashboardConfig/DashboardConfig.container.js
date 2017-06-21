//@flow
import { connect } from 'react-redux'
import DashboardConfig from './DashboardConfig'
import { selectedDashboardSelector } from '../selectors'
import { updateConfig } from '../actions'

const mapStateToProps = state => ({
  dashboard: selectedDashboardSelector(state),
})

const mapDispatchToProps = {
  changeName: updateConfig('name'),
  changeDescription: updateConfig('description'),
  changeCols: updateConfig('cols'),
  changeRows: updateConfig('rows'),
  changeRatio: updateConfig('ratio'),
}

export default connect(mapStateToProps, mapDispatchToProps)(DashboardConfig)
