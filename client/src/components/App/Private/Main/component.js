import React from 'react'
import { Route } from 'react-router-dom'

import Container from '@material-ui/core/Container'

import { makeUseStylesHook } from 'components/App/styles'

import Stats from './views/Stats'
import Channels from './views/Channels'

const Main = () => {
  const classes = makeUseStylesHook()()
  return (
    <main className={classes.content}>
      <div className={classes.appBarSpacer} />
      <Container maxWidth="lg" className={classes.container}>
        <Route path="/" exact component={Stats} />
        <Route path="/channels/" render={() => <Channels />} />
      </Container>
    </main>
  )
}

export default Main
