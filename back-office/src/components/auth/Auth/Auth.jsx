import React from 'react'
import LoginForm from '../LoginForm'

import './Auth.css'

const Auth = (props: PropsType) => {
  const { children, isLoggedIn } = props
  return <div className="Auth">{isLoggedIn ? children : <LoginForm />}</div>
}

export default Auth
