import React from 'react'
import ReactDOM from 'react-dom'
import App from './components/App'
import './css/index.css'

//Parse query parameters from URL and expose them in the location.queryParams variable
window.location.queryParams = {}
window.location.search
  .substr(1)
  .split('&')
  .forEach(function(pair) {
    if (pair === '') return
    var parts = pair.split('=')
    window.location.queryParams[parts[0]] =
      parts[1] && decodeURIComponent(parts[1].replace(/\+/g, ' '))
  })

ReactDOM.render(<App />, document.getElementById('root'))
