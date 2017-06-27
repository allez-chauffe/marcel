//@flow
import React from 'react'
import type { Children } from 'react'

import './Auth.css'

export type PropsType = {
  children: Children,
  isLoggedIn: boolean,
  login: () => void,
}

const Auth = (props: PropsType) => {
  const { children, isLoggedIn, login } = props
  return (
    <div className="Auth">
      {isLoggedIn
        ? children
        : <div>
            You are not logged In ! Please login before blabla
            <button onClick={login}>Login</button>
          </div>}
    </div>
  )
}

export default Auth
