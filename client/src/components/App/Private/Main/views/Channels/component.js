import React, { useState, useEffect } from 'react'
import clsx from 'clsx'

import Grid from '@material-ui/core/Grid'
import Paper from '@material-ui/core/Paper'

import channelsService from 'services/channelsService'

import { makeUseStylesHook } from 'components/App/styles'

import ChannelList from './ChannelList'
import SelectedChannel from './SelectedChannel'

const Channels = () => {
  const classes = makeUseStylesHook()()
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

  const minHeightPaper = clsx(classes.paper, classes.minHeightPaper)

  return (
    <Grid container spacing={3}>
      <Grid item xs={6}>
        <Paper className={minHeightPaper}>
          <ChannelList
            channels={channels}
            onDeleteChannel={handleDeleteChannel}
            onSelectChannel={setSelectedChannel}
          />
        </Paper>
      </Grid>
      <Grid item xs={6}>
        <Paper className={minHeightPaper}>
          <SelectedChannel
            selectedChannel={selectedChannel}
          />
        </Paper>
      </Grid>
    </Grid>
  )
}

export default Channels
