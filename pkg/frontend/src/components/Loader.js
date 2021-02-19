import React from 'react'
import '../css/loader.css'

const Loader = ({ className, style }) => (
  <div className={(className || '') + ' loaderContainer fullSize'}>
    <div className={'loader'} style={style}>
      Loading...
    </div>
  </div>
)

export default Loader
