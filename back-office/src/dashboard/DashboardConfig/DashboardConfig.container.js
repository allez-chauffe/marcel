//@flow
import { connect } from 'react-redux'
import DashboardConfig from './DashboardConfig'
import { dashboardSelector } from '../selectors'
import { updateConfig } from '../actions'

const mapStateToProps = state => ({
  dashboard: dashboardSelector(state),
})

const mapDispatchToProps = {
  changeName: updateConfig('name'),
  changeDescription: updateConfig('description'),
  changeCols: updateConfig('cols'),
  changeRows: updateConfig('rows'),
  changeRatio: updateConfig('ratio'),
}

export default connect(mapStateToProps, mapDispatchToProps)(DashboardConfig)
