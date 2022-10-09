import React from 'react'
import '../css/LoginForm.css'

const LoginForm = ({ login, password, onLoginChange, onPasswordChange, onLogin, onReset }) => (
  <div className="LoginForm">
    <form onSubmit={onLogin}>
      <div className="inputGroup">
        <label htmlFor="login">Login</label>
        <input id="login" value={login} onChange={onLoginChange} />
      </div>
      <div className="inputGroup">
        <label htmlFor="password">Mot de passe</label>
        <input id="password" type="password" value={password} onChange={onPasswordChange} />
      </div>
      <div className="actions">
        <button type="submit">Login</button>
        <button type="reset" onClick={onReset}>
          Annuler
        </button>
      </div>
    </form>
  </div>
)

export default LoginForm
