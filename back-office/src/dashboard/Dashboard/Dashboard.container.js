//@flow
import { connect } from 'react-redux'
import { dashboardSelector } from '../selectors'
import { uploadLayout } from '../actions'
import Dashboard from './Dashboard'

const mapStateToProps = state => ({
  dashboard: dashboardSelector(state),
})

const mapDispatchToProps = {
  uploadLayout,
}
