//@flow
import React from 'react'
import Input from 'react-toolbox/lib/input/Input'
import type { Dashboard } from '../type'

import './DashboardConfig.css'

export type PropsType = {
  dashboard: Dashboard,
}

const DashboardConfig = (props: PropsType) => {
  const {
    dashboard: { name, description, cols, rows, ratio },
    changeName,
    changeDescription,
    changeCols,
    changeRows,
    changeRatio,
  } = props

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
      <Input
        label="Ratio de l'écran"
        value={ratio}
        onChange={changeRatio}
        type="number"
      />
    </div>
  )
}

export default DashboardConfig
