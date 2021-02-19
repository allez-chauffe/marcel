import { connect } from 'react-redux'
import MediaConfig from './MediaConfig'
import { updateConfig, selectedDashboardSelector } from '../../../dashboard'

const mapStateToProps = state => {
  const dashboard = selectedDashboardSelector(state)
  if (!dashboard) throw new Error('A dashboard should be selected !')
  return {
    dashboard,
  }
}

const mapDispatchToProps = {
  changeName: updateConfig('name'),
  changeDescription: updateConfig('description'),
  changeCols: updateConfig('cols'),
  changeRows: updateConfig('rows'),
  changeRatio: updateConfig('screenRatio'),
  changeDisplayGrid: updateConfig('displayGrid'),
  changeBackgroundColor: updateConfig('stylesvar.background-color'),
  changePrimaryColor: updateConfig('stylesvar.primary-color'),
  changeSecondaryColor: updateConfig('stylesvar.secondary-color'),
  changeFontFamily: updateConfig('stylesvar.font-family'),
}

export default connect(
  mapStateToProps,
  mapDispatchToProps,
)(MediaConfig)
