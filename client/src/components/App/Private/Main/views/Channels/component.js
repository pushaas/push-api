import React, { useContext, useEffect, useState } from 'react'
import { Redirect } from 'react-router-dom'
import clsx from 'clsx'

import Grid from '@material-ui/core/Grid'
import Paper from '@material-ui/core/Paper'

import { privateChannelsPath } from 'navigation'

import channelsService from 'services/channelsService'

import SetTitleContext from 'components/contexts/SetTitleContext'
import { useStyles } from 'components/App/Private/styles'

import ChannelList from './ChannelList'
import SelectedChannel from './SelectedChannel'

const Channels = (props) => {
  const id = props.match.params.id
  const classes = useStyles()
  const [didLoad, setDidLoad] = useState(false)
  const [channels, setChannels] = useState([])
  const setTitle = useContext(SetTitleContext)

  const findSelectedChannelById = () => {
    if (id && channels.length) {
      return channels.find(c => c.id === id)
    }
  }
  const selectedChannel = id ? findSelectedChannelById() : undefined

  useEffect(() => {
    setTitle('Persistent Channels')
  }, [setTitle])

  useEffect(() => {
    channelsService.getChannels()
      .then((data) => {
        setChannels(data)
        setDidLoad(true)
      })
  }, [])

  const handleDeleteChannel = (channel) => {
    channelsService.deleteChannel(channel.id)
      .then(() => setChannels(channels.filter(c => c !== channel)))
  }

  const channelsMinHeightPaper = clsx(classes.paper, classes.channelsMinHeightPaper)

  if (didLoad && id && !selectedChannel) {
    return (
      <Redirect to={privateChannelsPath} />
    )
  }

  return (
    <Grid container spacing={3}>
      <Grid item xs={6}>
        <Paper className={channelsMinHeightPaper}>
          <ChannelList
            channels={channels}
            onDeleteChannel={handleDeleteChannel}
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
