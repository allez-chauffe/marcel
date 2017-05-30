// @flow
import React, { Component } from 'react'
import { Provider } from 'react-redux'
import AppBar from 'react-toolbox/lib/app_bar/AppBar'

import '../assets/react-toolbox/theme.css'
import ThemeProvider from 'react-toolbox/lib/ThemeProvider'
import theme from '../assets/react-toolbox/theme.js'

import MainScreen from '../MainScreen'

import store from '../store'

import './App.css'

export default class App extends Component {
  render() {
    return (
      <ThemeProvider theme={theme}>
        <Provider store={store}>
          <div className="App">
            <header>
              <AppBar title="Zenboard" leftIcon="menu" />
            </header>
            <main>
              <MainScreen />
            </main>
            <footer>
              <AppBar />
            </footer>
          </div>
        </Provider>
      </ThemeProvider>
    )
  }
}
