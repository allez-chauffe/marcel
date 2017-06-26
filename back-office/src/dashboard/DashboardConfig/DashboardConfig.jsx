//@flow
import React from 'react'
import { find } from 'lodash'
import Input from 'react-toolbox/lib/input/Input'
import Dropdown from 'react-toolbox/lib/dropdown/Dropdown'
import type { Dashboard } from '../type'

import './DashboardConfig.css'

export type PropsType = {
  dashboard: Dashboard,
  changeName: string => void,
  changeDescription: string => void,
  changeCols: number => void,
  changeRows: number => void,
  changeRatio: number => void,
}

const DashboardConfig = (props: PropsType) => {
  const { dashboard } = props

  const {
    changeName,
    changeDescription,
    changeCols,
    changeRows,
    changeRatio,
  } = props
  const { name, description, cols, rows, ratio } = dashboard
  // const ratio = find(ratios, { value: rawRatio })

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
    </div>
  )
}

export default DashboardConfig
