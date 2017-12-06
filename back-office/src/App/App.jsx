// @flow
import React, { Component } from 'react'
import { Provider } from 'react-redux'
import ThemeProvider from 'react-toolbox/lib/ThemeProvider'
import theme from '../assets/react-toolbox/theme.js'

import '../assets/react-toolbox/theme.css'
import 'react-redux-toastr/lib/css/react-redux-toastr.min.css'

import store from '../store'
import { AppLayout } from '../pages'

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
