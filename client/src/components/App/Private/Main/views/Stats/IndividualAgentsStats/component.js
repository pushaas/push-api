import React from 'react'
import clsx from 'clsx'

import Grid from '@material-ui/core/Grid'
import GridList from '@material-ui/core/GridList';
import GridListTile from '@material-ui/core/GridListTile';
import Paper from '@material-ui/core/Paper'
import Typography from '@material-ui/core/Typography'

import dateService from 'services/dateService'

import Title from 'components/common/Title'

const renderTextItem = (label, value) => (
  <div>
    <Typography component="span" display="inline" variant="overline" color="primary" >{label} </Typography>
    <Typography component="span" display="inline" variant="body2">{value}</Typography>
  </div>
)

const IndividualAgentsStats = ({ stats, classes }) => {
  const statsMinHeightPaper = clsx(classes.paper, classes.statsMinHeightPaper)

  return (
    <React.Fragment>
      {stats ? (stats.all.map(a => (
        <Grid key={a.agent} item xs={12}>
          <Paper className={statsMinHeightPaper}>
            <GridList cellHeight={'auto'} cols={2}>
              <GridListTile cols={1}>
                <Title>Agent {a.agent}</Title>
                {renderTextItem('agent', a.agent)}
                {renderTextItem('hostname', a.hostname)}
                {renderTextItem('subscribers', a.subscribers)}
                {renderTextItem('uptime', a.uptime)}
                {renderTextItem('updated on', dateService.formatDate(a.time))}
              </GridListTile>
              <GridListTile cols={1}>
                {renderTextItem('channels', a.channels)}
                {renderTextItem('channels in delete', a.channels_in_delete)}
                {renderTextItem('channels in trash', a.channels_in_trash)}
                {renderTextItem('wildcard channels', a.wildcard_channels)}
                {renderTextItem('published messages', a.published_messages)}
                {renderTextItem('stored messages', a.stored_messages)}
                {renderTextItem('messages in trash', a.messages_in_trash)}
              </GridListTile>
            </GridList>
          </Paper>
        </Grid>
      ))) : null}
    </React.Fragment>
  )
}

export default IndividualAgentsStats
