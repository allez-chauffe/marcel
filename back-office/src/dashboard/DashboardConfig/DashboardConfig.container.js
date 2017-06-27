//@flow
import { connect } from 'react-redux'
import DashboardConfig from './DashboardConfig'
import { selectedDashboardSelector } from '../selectors'
import { updateConfig } from '../actions'

const mapStateToProps = state => {
  const dashboard = selectedDashboardSelector(state)
  if (!dashboard) throw new Error('A dashboard should be selected !')
  return { dashboard }
}

const mapDispatchToProps = {
  changeName: updateConfig('name'),
  changeDescription: updateConfig('description'),
  changeCols: updateConfig('cols'),
  changeRows: updateConfig('rows'),
  changeRatio: updateConfig('ratio'),
}

export default connect(mapStateToProps, mapDispatchToProps)(DashboardConfig)
