//@flow
import React, { Component } from 'react'
import type { Children } from 'react'
import ProgressBar from 'react-toolbox/lib/progress_bar/ProgressBar'
import LoginForm from '../LoginForm'

import './Auth.css'

export type PropsType = {
  children: Children,
  isLoggedIn: boolean,
  isLoading: boolean,
  login: () => void,
}

class Auth extends Component {
  props: PropsType

  componentWillMount() {
    const { isLoggedIn, isLoading, login } = this.props
    if (!isLoggedIn && !isLoading) login()
  }

  render() {
    const { children, isLoggedIn, isLoading } = this.props

    if (isLoading) return <ProgressBar type="circular" mode="indeterminate" className="loader" />

    return (
      <div className="Auth">
        {isLoading ? (
          <ProgressBar type="circular" mode="indeterminate" className="loader" />
        ) : isLoggedIn ? (
          children
        ) : (
          <LoginForm />
        )}
      </div>
    )
  }
}

export default Auth
