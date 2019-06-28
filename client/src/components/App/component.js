import React from 'react'
import { BrowserRouter as Router } from 'react-router-dom'
import CssBaseline from '@material-ui/core/CssBaseline'

import Private from './Private'

// import channelsService from 'services/channelsService'
// import statsService from 'services/statsService'

// componentDidMount() {
//   channelsService.getChannels()
//     .then((data) => {
//       console.log('### getChannels data', data)
//     })
//
//   statsService.getGlobalStats()
//     .then((data) => {
//       console.log('### getGlobalStats data', data)
//     })
// }

const App = () => (
  <React.Fragment>
    <CssBaseline />
    <Router basename="/admin">
      <Private />
    </Router>
  </React.Fragment>
)

export default App
