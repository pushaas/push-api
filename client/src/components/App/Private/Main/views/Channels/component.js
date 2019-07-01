import React, { useState, useEffect } from 'react'
import clsx from 'clsx'

import Grid from '@material-ui/core/Grid'
import Paper from '@material-ui/core/Paper'

import channelsService from 'services/channelsService'

import { useStyles } from 'components/App/Private/styles'

import ChannelList from './ChannelList'
import SelectedChannel from './SelectedChannel'

const Channels = () => {
  const classes = useStyles()
  const [selectedChannel, setSelectedChannel] = useState(null)
  const [channels, setChannels] = useState([])

  useEffect(() => {
    channelsService.getChannels()
      .then((data) => {
        setChannels(data)
        if (data.length) {
          setSelectedChannel(data[0])
        }
      })
  }, [])

  const handleDeleteChannel = (channel) => {
    channelsService.deleteChannel(channel.id)
      .then(() => {
        if (channel === selectedChannel) {
          setSelectedChannel(null)
        }
        setChannels(channels.filter(c => c !== channel))
      })
  }

  const channelsMinHeightPaper = clsx(classes.paper, classes.channelsMinHeightPaper)

  return (
    <Grid container spacing={3}>
      <Grid item xs={6}>
        <Paper className={channelsMinHeightPaper}>
          <ChannelList
            channels={channels}
            onDeleteChannel={handleDeleteChannel}
            onSelectChannel={setSelectedChannel}
          />
        </Paper>
      </Grid>
      <Grid item xs={6}>
        <Paper className={channelsMinHeightPaper}>
          <SelectedChannel channel={selectedChannel} />
        </Paper>
      </Grid>
    </Grid>
  )
}

export default Channels
