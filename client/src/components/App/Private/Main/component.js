import React from 'react'
import Container from '@material-ui/core/Container'
// import Grid from '@material-ui/core/Grid'
// import Paper from '@material-ui/core/Paper'

import { Route } from 'react-router-dom'

import Stats from './views/Stats'
import Channels from './views/Channels'

// {/* <Grid container spacing={3}>
//   {/* Chart */}
//   <Grid item xs={12} md={8} lg={9}>
//     <Paper className={fixedHeightPaper}>
//       {/* <Chart /> */}
//     </Paper>
//   </Grid>
//   {/* Recent Deposits */}
//   <Grid item xs={12} md={4} lg={3}>
//     <Paper className={fixedHeightPaper}>
//       {/* <Deposits /> */}
//     </Paper>
//   </Grid>
//   {/* Recent Orders */}
//   <Grid item xs={12}>
//     <Paper className={classes.paper}>
//       {/* <Orders /> */}
//     </Paper>
//   </Grid>
// </Grid> */}

const Main = ({ classes, fixedHeightPaper }) => (
  <main className={classes.content}>
    <div className={classes.appBarSpacer} />
    <Container maxWidth="lg" className={classes.container}>
      <Route path="/" exact component={Stats} />
      <Route path="/channels/" component={Channels} />
    </Container>
  </main>
)

export default Main
