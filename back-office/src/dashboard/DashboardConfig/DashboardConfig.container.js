//@flow
import { connect } from 'react-redux'
import DashboardConfig from './DashboardConfig'
import { selectedDashboardSelector, displayGridSelector } from '../selectors'
import { updateConfig, toggleDisplayGrid } from '../actions'

const mapStateToProps = state => {
  const dashboard = selectedDashboardSelector(state)
  if (!dashboard) throw new Error('A dashboard should be selected !')
  return {
    dashboard,
    displayGrid: displayGridSelector(state),
  }
}

const mapDispatchToProps = {
  changeName: updateConfig('name'),
  changeDescription: updateConfig('description'),
  changeCols: updateConfig('cols'),
  changeRows: updateConfig('rows'),
  changeRatio: updateConfig('ratio'),
  changeBackgroundColor: updateConfig('stylesvar.background-color'),
  changePrimaryColor: updateConfig('stylesvar.primary-color'),
  changeSecondaryColor: updateConfig('stylesvar.secondary-color'),
  changeFontFamily: updateConfig('stylesvar.font-family'),
  toggleDisplayGrid,
}

export default connect(mapStateToProps, mapDispatchToProps)(DashboardConfig)
