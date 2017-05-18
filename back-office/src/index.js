import React from 'react'
import ReactDOM from 'react-dom'
import App from './App'
import './index.css'
import '../assets/react-toolbox/theme.css'
import ThemeProvider from 'react-toolbox/lib/ThemeProvider'
import theme from '../assets/react-toolbox/theme.js';

ReactDOM.render(
    <ThemeProvider theme={theme}>
        <App />
    </ ThemeProvider>,
    document.getElementById('root'),
)
