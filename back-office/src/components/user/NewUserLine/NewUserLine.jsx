import React, { Component } from 'react'
import Input from 'react-toolbox/lib/input/Input'
import Dropdown from 'react-toolbox/lib/dropdown/Dropdown'
import Button from 'react-toolbox/lib/button/Button'

import './NewUserLine.css'

class NewUserLine extends Component {
  state = {
    passwordConfirmError: false,
  }

  roles = [{ value: 'user', label: 'user' }, { value: 'admin', label: 'admin' }]

  handleSubmit = e => {
    e.preventDefault()
    const user = this.props.user
    if (user.password !== user.confirmPassword) {
      this.setState({ passwordConfirmError: true })
      return
    } else {
      this.setState({ passwordConfirmError: false })
    }

    this.props.handleSave(user)
  }

  render() {
    const { user, handleChange, handleReset, disableInputRole } = this.props

    let buttonReset = ''
    if (handleReset) {
      buttonReset = <Button icon="not_interested" label="Reset" raised onClick={handleReset} />
    }

    let inputRole = ''
    if (!disableInputRole) {
      inputRole = (
        <Dropdown
          label="Rôle"
          source={this.roles}
          onChange={handleChange.bind(this, 'role')}
          value={user.role}
          required
        />
      )
    }

    return (
      <form onSubmit={e => this.handleSubmit(e)} className="UserForm">
        <Input
          type="text"
          label="Nom"
          name="fullname"
          value={user.displayName}
          onChange={handleChange.bind(this, 'displayName')}
          required
        />
        <Input
          type="text"
          label="Login"
          name="login"
          value={user.login}
          onChange={handleChange.bind(this, 'login')}
          required
          maxLength={16}
        />
        {inputRole}
        <Input
          type="password"
          label="Mot de passe"
          name="password"
          value={user.password}
          onChange={handleChange.bind(this, 'password')}
        />
        <Input
          type="password"
          label="Confirmation mot de passe"
          error={this.state.passwordConfirmError ? 'Password confirmation error' : ''}
          name="confirmPassword"
          value={user.confirmPassword}
          onChange={handleChange.bind(this, 'confirmPassword')}
        />
        <Button
          type="submit"
          icon="save"
          label={user.id ? 'Sauvegarder' : 'Créer'}
          raised
          primary
        />
        {buttonReset}
      </form>
    )
  }
}

export default NewUserLine
