// @flow
import React from 'react'
import Input from 'react-toolbox/lib/input/Input'

import './SearchField.css'

const SearchField = (props: {
  label: string,
  value: string,
  onChange: string => void,
}) => {
  const { label, value, onChange } = props
  return <Input label={label} icon="search" value={value} onChange={onChange} />
}

export default SearchField
