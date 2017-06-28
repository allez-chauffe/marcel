//@flow
import React from 'react'
import type { Children } from 'react'
import LoginForm from '../LoginForm'

import './Auth.css'

export type PropsType = {
  children: Children,
  isLoggedIn: boolean,
}

const Auth = (props: PropsType) => {
  const { children, isLoggedIn } = props
  return (
    <div className="Auth">
      {isLoggedIn ? children : <LoginForm />}
    </div>
  )
}

export default Auth
