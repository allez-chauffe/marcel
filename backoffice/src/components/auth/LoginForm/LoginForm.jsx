import React, { Component } from 'react'
import Card from 'react-toolbox/lib/card/Card'
import CardTitle from 'react-toolbox/lib/card/CardTitle'
import Input from 'react-toolbox/lib/input/Input'
import Button from 'react-toolbox/lib/button/Button'

import './LoginForm.css'

class LoginForm extends Component {
  onKeyDown = event => {
    const enterKeyCode = 13
    if (event.keyCode === enterKeyCode || event.which === enterKeyCode) this.props.loginAction()
  }

  render() {
    const { login, password } = this.props
    const { loginAction, changeLogin, changePassword, reset } = this.props
    return (
      <Card className="LoginForm" style={{ width: '25em' }}>
        <form>
          <CardTitle>Authentification</CardTitle>
          <Input label="Login" value={login} onChange={changeLogin} onKeyDown={this.onKeyDown} />
          <Input
            label="Mot de passe"
            type="password"
            value={password}
            onChange={changePassword}
            onKeyDown={this.onKeyDown}
          />
          <div className="actions">
            <Button onClick={loginAction}>Login</Button>
            <Button onClick={reset}> Reset </Button>
          </div>
        </form>
      </Card>
    )
  }
}

export default LoginForm
