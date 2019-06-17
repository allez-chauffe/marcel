import React from 'react'
import Input from 'react-toolbox/lib/input/Input'

import './SearchField.css'

const SearchField = props => {
  const { label, value, onChange } = props
  return <Input label={label} icon="search" value={value} onChange={onChange} />
}

export default SearchField
