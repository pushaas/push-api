import React, { useState, useEffect } from 'react'
import clsx from 'clsx'

import Grid from '@material-ui/core/Grid'
import Paper from '@material-ui/core/Paper'
import Typography from '@material-ui/core/Typography'

import dateService from 'services/dateService'
import statsService from 'services/statsService'

import { useStyles } from 'components/App/Private/styles'
import Title from 'components/common/Title'
import IndividualAgentsStats from './IndividualAgentsStats'

const AggregatedAgentsStats = ({ stats, classes }) => (
  <React.Fragment>
    <Title>Agents</Title>
    {stats ? (
      <Typography component="p" variant="h4">
        {stats.all.length}
      </Typography>
    ) : (null)}
    {stats ? (
      <Typography color="textSecondary" className={classes.statsSubscribers}>
        on {dateService.formatDate(stats.aggregated.time)}
      </Typography>
    ) : (null)}
  </React.Fragment>
)

const SubscribersStats = ({ stats, classes }) => (
  <React.Fragment>
    <Title>Subscribers</Title>
    {stats ? (
      <Typography component="p" variant="h4">
        {stats.aggregated.subscribers}
      </Typography>
    ) : (null)}
    {stats ? (
      <Typography color="textSecondary" className={classes.statsSubscribers}>
        on {dateService.formatDate(stats.aggregated.time)}
      </Typography>
    ) : (null)}
  </React.Fragment>
)

const Stats = () => {
  const classes = useStyles()
  const [stats, setStats] = useState()

  useEffect(() => {
    statsService.getGlobalStats()
      .then((data) => {
        setStats(data)
      })
  }, [])

  const statsMinHeightPaper = clsx(classes.paper, classes.statsMinHeightPaper)

  return (
    <React.Fragment>
      <Grid container spacing={3}>
        <Grid item xs={12} md={6}>
          <Paper className={statsMinHeightPaper}>
            <AggregatedAgentsStats classes={classes} stats={stats} />
          </Paper>
        </Grid>
        <Grid item xs={12} md={6}>
          <Paper className={statsMinHeightPaper}>
            <SubscribersStats classes={classes} stats={stats} />
          </Paper>
        </Grid>
        <IndividualAgentsStats classes={classes} stats={stats} />
      </Grid>
    </React.Fragment>
  )
}

export default Stats
