import React, { Component } from 'react'
import { BrowserRouter as Router } from 'react-router-dom'
import CssBaseline from '@material-ui/core/CssBaseline'

import channelsService from 'services/channelsService'
import statsService from 'services/statsService'

import Private from './Private'

class App extends Component {
  componentDidMount() {
    channelsService.getChannels()
      .then((data) => {
        console.log('### getChannels data', data)
      })

    statsService.getGlobalStats()
      .then((data) => {
        console.log('### getGlobalStats data', data)
      })
  }

  render() {
    return (
      <React.Fragment>
        <CssBaseline />
        <Router basename="/admin">
          <Private />
        </Router>
      </React.Fragment>
    )
  }
}

export default App
