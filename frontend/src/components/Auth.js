import React, { Children } from 'react'
import { toast } from 'react-toastify'
import { authFetcher } from '../utils/fetcher'
import Loader from './Loader'
import LoginForm from './LoginForm'

class Auth extends React.Component {
  state = {
    user: null,
    loading: true,
    form: {
      login: '',
      password: '',
    },
  }

  componentDidMount() {
    this.authFetcher = authFetcher(this.props.config)
    this.login()
  }

  login = event => {
    event && event.preventDefault()
    const { login, password } = this.state.form
    console.log('Logging in...')
    this.authFetcher
      .post('/login', login && password ? { login, password } : null)
      .then(user => this.setState({ user, loading: false }))
      .catch(res => {
        if (res.status !== 403) toast.error("Impossible de contacter le serveur d'authentification")
        else if (login || password) toast.error('Mauvais login ou mot de passe')
        this.setState({ loading: false })
      })
    return false
  }

  reset = () => {
    this.setState({ form: { login: '', password: '' } })
  }

  loginChange = event => {
    this.setState({ form: { ...this.state.form, login: event.target.value } })
  }

  passwordChange = event => {
    console.log(event.target.value)
    this.setState({ form: { ...this.state.form, password: event.target.value } })
  }

  render() {
    const { user, loading, form } = this.state

    if (loading) return <Loader />

    if (user)
      return Children.map(this.props.children, child =>
        React.cloneElement(child, { config: { ...this.props.config, user } }),
      )

    return (
      <LoginForm
        login={form.login}
        password={form.password}
        onLoginChange={this.loginChange}
        onPasswordChange={this.passwordChange}
        onLogin={this.login}
        onReset={this.reset}
      />
    )
  }
}

export default Auth
