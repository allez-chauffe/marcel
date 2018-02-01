//@flow
import React from 'react'
import Input from 'react-toolbox/lib/input/Input'
import Dropdown from 'react-toolbox/lib/dropdown/Dropdown'
import Switch from 'react-toolbox/lib/switch/Switch'
import { ColorPicker } from '../../../common'
import type { Dashboard } from '../../../dashboard/type'

import './MediaConfig.css'

export type PropsType = {
  dashboard: Dashboard,
  changeName: string => void,
  changeDescription: string => void,
  changeCols: number => void,
  changeRows: number => void,
  changeRatio: number => void,
  changeDisplayGrid: boolean => void,
  changeBackgroundColor: string => void,
  changePrimaryColor: string => void,
  changeSecondaryColor: string => void,
  changeFontFamily: string => void,
}

const MediaConfig = (props: PropsType) => {
  const { dashboard } = props

  const {
    changeName,
    changeDescription,
    changeCols,
    changeRows,
    changeRatio,
    changeDisplayGrid,
    changeBackgroundColor,
    changePrimaryColor,
    changeSecondaryColor,
    changeFontFamily,
  } = props
  const { name, description, cols, rows, screenRatio, displayGrid } = dashboard

  return (
    <div className="MediaConfig">
      <Input label="Nom" value={name} onChange={changeName} />
      <Input label="Description" value={description} onChange={changeDescription} multiline />
      <Input label="Nombre de colonnes" value={cols} type="number" onChange={changeCols} />
      <Input label="Nombre de lignes" value={rows} onChange={changeRows} type="number" />
      <Dropdown
        source={[
          { label: '16/9', value: 16 / 9 },
          { label: '16/9 (portrait)', value: 9 / 16 },
          { label: '4/3', value: 4 / 3 },
          { label: '4/3 (protrait)', value: 3 / 4 },
        ]}
        value={screenRatio}
        onChange={changeRatio}
      />
      <ColorPicker
        value={dashboard.stylesvar['background-color']}
        onChange={changeBackgroundColor}
        label="Background color"
      />
      <ColorPicker
        value={dashboard.stylesvar['primary-color']}
        onChange={changePrimaryColor}
        label="Primary color"
      />
      <ColorPicker
        value={dashboard.stylesvar['secondary-color']}
        onChange={changeSecondaryColor}
        label="Secondary color"
      />
      <Input
        label="Font family"
        value={dashboard.stylesvar['font-family']}
        onChange={changeFontFamily}
      />
      <div className="gridDisplay">
        <Switch label="Afficher la grille" checked={displayGrid} onChange={changeDisplayGrid} />
      </div>
    </div>
  )
}

export default MediaConfig
