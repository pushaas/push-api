import React from 'react'
import { BrowserRouter as Router } from 'react-router-dom'
import CssBaseline from '@material-ui/core/CssBaseline'

import { routerBaseName } from 'navigation'

import Private from './Private'

const App = () => (
  <React.Fragment>
    <CssBaseline />
    <Router basename={routerBaseName}>
      <Private />
    </Router>
  </React.Fragment>
)

export default App
