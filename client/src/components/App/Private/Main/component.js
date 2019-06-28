import React from 'react'
import Container from '@material-ui/core/Container'

import { Route } from 'react-router-dom'

import Stats from './views/Stats'
import Channels from './views/Channels'

const Main = ({ classes, fixedHeightPaper }) => (
  <main className={classes.content}>
    <div className={classes.appBarSpacer} />
    <Container maxWidth="lg" className={classes.container}>
      <Route path="/" exact component={Stats} />
      <Route path="/channels/" render={() => <Channels classes={classes} fixedHeightPaper={fixedHeightPaper} />} />
    </Container>
  </main>
)

export default Main
