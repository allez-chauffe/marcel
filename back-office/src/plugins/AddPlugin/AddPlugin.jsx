import React, { useState } from 'react'
import Input from 'react-toolbox/lib/input/Input'
import './AddPlugin.css'
import Button from 'react-toolbox/lib/button'
import FontIcon from 'react-toolbox/lib/font_icon/FontIcon'
import { LoadingIndicator } from '../../components/commons'

const AddPlugin = ({ add, adding }) => {
  const [url, setUrl] = useState('')
  const onSubmit = event => {
    event.preventDefault()
    add(url)
  }

  return (
    <form className="AddPlugin" onSubmit={onSubmit}>
      <Input
        disabled={adding}
        label="Ajouter un plugin"
        onChange={value => setUrl(value)}
        value={url}
        hint="https://github.com/Zenika/marcel-plugin-helloworld"
        name="pluginUrl"
        className="UrlInput"
      />
      {adding ? (
        <LoadingIndicator />
      ) : (
        <Button primary raised onClick={() => add(url)} disabled={!url}>
          <FontIcon>add</FontIcon> Ajouter
        </Button>
      )}
    </form>
  )
}

export default AddPlugin
