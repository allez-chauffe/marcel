import React from 'react'
import LoginForm from '../LoginForm'

import './Auth.css'

const Auth = (props: PropsType) => {
  const { children, isLoggedIn } = props
  if (isLoggedIn) return children
  return (
    <div className="Auth">
      <LoginForm />
    </div>
  )
}

export default Auth
