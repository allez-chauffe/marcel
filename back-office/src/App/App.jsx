// @flow
import React, { Component } from 'react'
import { Provider } from 'react-redux'

import '../assets/react-toolbox/theme.css'
import ThemeProvider from 'react-toolbox/lib/ThemeProvider'
import theme from '../assets/react-toolbox/theme.js'

import store from '../store'
import { AppLayout } from '../layouts'

export default class App extends Component {
  render() {
    return (
      <ThemeProvider theme={theme}>
        <Provider store={store}>
          <AppLayout />
        </Provider>
      </ThemeProvider>
    )
  }
}
