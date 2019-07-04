import React, { useContext, useState, useEffect } from 'react'
import Divider from '@material-ui/core/Divider'

import dateService from 'services/dateService'
import statsService from 'services/statsService'

import ConfigContext from 'components/contexts/ConfigContext'
import { useStyles } from 'components/App/Private/styles'
import Title from 'components/common/Title'
import {
  renderSubtitle1,
  renderSubtitle2,
  renderTextItem,
} from 'components/common/textHelpers'
import Messages from './Messages'

const NoSelectedChannel = () => renderTextItem('No channel selected')

const SelectedChannelInfo = ({ channel }) => (
  <React.Fragment>
    {renderSubtitle1('Channel Info')}
    {renderTextItem('id', channel.id)}
    {renderTextItem('ttl', channel.ttl)}
    {renderTextItem('created', dateService.formatDate(channel.created))}
    {renderTextItem('expiration', channel.ttl ? (dateService.calculateExpiration(channel.created, channel.ttl)) : 'never')}
  </React.Fragment>
)

const SelectedChannelStats = ({ channel }) => {
  const [stats, setStats] = useState(null)

  useEffect(() => {
    statsService.getChannelStats(channel.id)
      .then((data) => {
        setStats(data)
      })
  }, [channel.id])

  if (!stats) {
    return null
  }

  return (
    <React.Fragment>
      {renderSubtitle1('Channel Stats')}
      {renderSubtitle2('Aggregated')}
      {renderTextItem('subscribers', stats.aggregated.subscribers)}
      {renderSubtitle2('By agent')}
      {stats.all.map((a, i) => (
        <React.Fragment key={a.agent}>
          {i > 0 ? (<Divider light />) : null}
          {renderTextItem('agent', a.agent)}
          {renderTextItem('push-stream hostname', a.hostname)}
          {renderTextItem('published messages', a.published_messages)}
          {renderTextItem('stored messages', a.stored_messages)}
          {renderTextItem('subscribers', a.subscribers)}
        </React.Fragment>
      ))}
    </React.Fragment>
  )
}

const SelectedChannelDetails = ({ channel, classes }) => {
  const config = useContext(ConfigContext)
  if (!config) {
    return null
  }

  return (
    <React.Fragment>
      <Divider className={classes.channelInfoDivider} />
      <SelectedChannelInfo channel={channel} />
      <Divider className={classes.channelInfoDivider} />
      <Messages channel={channel} />
      <Divider className={classes.channelInfoDivider} />
      <SelectedChannelStats channel={channel} />
    </React.Fragment>
  )
}

const SelectedChannel = ({ channel }) => {
  const classes = useStyles()

  return (
    <React.Fragment>
      <Title>
        Selected Channel
      </Title>
      {channel ? (<SelectedChannelDetails channel={channel} classes={classes} />) : (<NoSelectedChannel />)}
    </React.Fragment>
  )
}

export default SelectedChannel
