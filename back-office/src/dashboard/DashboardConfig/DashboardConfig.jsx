//@flow
import React from 'react'
import Input from 'react-toolbox/lib/input/Input'
import Dropdown from 'react-toolbox/lib/dropdown/Dropdown'
import Switch from 'react-toolbox/lib/switch/Switch'
import { ColorPicker } from '../../common'
import type { Dashboard } from '../type'

import './DashboardConfig.css'

export type PropsType = {
  dashboard: Dashboard,
  displayGrid: boolean,
  changeName: string => void,
  changeDescription: string => void,
  changeCols: number => void,
  changeRows: number => void,
  changeRatio: number => void,
  toggleDisplayGrid: () => void,
}

const DashboardConfig = (props: PropsType) => {
  const { dashboard, displayGrid } = props

  const {
    changeName,
    changeDescription,
    changeCols,
    changeRows,
    changeRatio,
    toggleDisplayGrid,
  } = props
  const { name, description, cols, rows, ratio } = dashboard

  return (
    <div className="DashboardConfig">
      <Input label="Nom" value={name} onChange={changeName} />
      <Input
        label="Description"
        value={description}
        onChange={changeDescription}
        multiline
      />
      <Input
        label="Nombre de colonnes"
        value={cols}
        type="number"
        onChange={changeCols}
      />
      <Input
        label="Nombre de lignes"
        value={rows}
        onChange={changeRows}
        type="number"
      />
      <Dropdown
        source={[
          { label: '16/9', value: 16 / 9 },
          { label: '16/9 (portrait)', value: 9 / 16 },
          { label: '4/3', value: 4 / 3 },
          { label: '4/3 (protrait)', value: 3 / 4 },
        ]}
        value={ratio}
        onChange={changeRatio}
      />
      <ColorPicker
        value="#700"
        onChange={console.log.bind(console)}
        label="Background color"
      />
      <Switch
        style={{ fontSize: '16px' }}
        label="Afficher la grille"
        checked={displayGrid}
        onChange={toggleDisplayGrid}
      />
    </div>
  )
}

export default DashboardConfig
