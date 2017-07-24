//@flow
import { flow } from 'lodash'
import { connect } from 'react-redux'
import { SortableElement } from 'react-sortable-hoc'
import { selectPlugin } from '../../../../dashboard'
import SortablePlugin from './SortablePlugin'

const hoc = flow(connect(null, { selectPlugin }), SortableElement)

export default hoc(SortablePlugin)
