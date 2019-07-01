import React, { useState } from 'react'
import { BrowserRouter as Router } from 'react-router-dom'
import CssBaseline from '@material-ui/core/CssBaseline'

import { routerBaseName } from 'navigation'

import { UserContext } from 'components/contexts/UserContext'
import SetUserContext from 'components/contexts/SetUserContext'
import Private from './Private'
import Public from './Public'

const App = () => {
  const [user, setUser] = useState(null)
  return (
    <React.Fragment>
      <CssBaseline />
      <Router basename={routerBaseName}>
        <UserContext.Provider value={user}>
          <SetUserContext.Provider value={setUser}>
            {user ? (<Private />) : (<Public />)}
          </SetUserContext.Provider>
        </UserContext.Provider>
      </Router>
    </React.Fragment>
  )
}

export default App
