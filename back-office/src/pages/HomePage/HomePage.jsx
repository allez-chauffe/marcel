//@flow
import React, { Component } from 'react'

class HomePage extends Component {
  props: {
    goToMedias: () => {},
  }

  componentDidMount() {
    this.props.goToMedias()
  }

  render() {
    return <div>Coucou</div>
  }
}

export default HomePage
