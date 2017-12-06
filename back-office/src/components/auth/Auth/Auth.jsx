//@flow
import React from 'react'
import type { Children } from 'react'
import LoginForm from '../LoginForm'

import './Auth.css'

export type PropsType = {
  children: Children,
  isLoggedIn: boolean,
}

const Auth = (props: PropsType) => (
  <div className="Auth">{props.isLoggedIn ? props.children : <LoginForm />}</div>
)

export default Auth
