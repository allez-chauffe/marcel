import { flow } from 'lodash'
import { connect } from 'react-redux'
import { SortableElement } from 'react-sortable-hoc'
import { selectPlugin, deletePlugin } from '../../../../dashboard'
import SortablePlugin from './SortablePlugin'

const mapDispatchToProps = (dispatch, ownProps) => ({
  onDelete: () => dispatch(deletePlugin(ownProps.plugin)),
  onSelect: () => dispatch(selectPlugin(ownProps.plugin)),
})

const hoc = flow(
  connect(
    null,
    mapDispatchToProps,
  ),
  SortableElement,
)

export default hoc(SortablePlugin)
