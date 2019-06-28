import React, { useState, useEffect } from 'react'

import Grid from '@material-ui/core/Grid'
import Paper from '@material-ui/core/Paper'

import channelsService from 'services/channelsService'

import ChannelList from './ChannelList'
import SelectedChannel from './SelectedChannel'

const Channels = ({classes, fixedHeightPaper}) => {
  const [selectedChannel, setSelectdChannel] = useState()
  const [channels, setChannels] = useState([])

  useEffect(() => {
    channelsService.getChannels()
      .then((data) => {
        setChannels(data)
      })
  }, [])

  return (
    <Grid container spacing={3}>
      <Grid item xs={4}>
        <Paper className={fixedHeightPaper}>
          <ChannelList channels={channels} onSelectChannel={setSelectdChannel} />
        </Paper>
      </Grid>
      <Grid item xs={8}>
        <Paper className={fixedHeightPaper}>
          <SelectedChannel selectedChannel={selectedChannel} />
        </Paper>
      </Grid>
    </Grid>
  )
}

export default Channels
