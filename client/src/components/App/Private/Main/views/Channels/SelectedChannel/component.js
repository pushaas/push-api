import React, { useState, useEffect } from 'react'
import Divider from '@material-ui/core/Divider'
import Typography from '@material-ui/core/Typography'

import dateService from 'services/dateService'
import statsService from 'services/statsService'

import { makeUseStylesHook } from 'components/App/styles'
import Title from 'components/common/Title'

const renderSubtitle1 = (text) => (<Typography component="p" variant="subtitle1" color="primary" gutterBottom>{text}</Typography>)
const renderSubtitle2 = (text) => (<Typography component="p" variant="subtitle2" color="primary" gutterBottom>{text}</Typography>)
const renderTextItem = (label, value) => (
  <div>
    <Typography component="span" display="inline" variant="overline" color="primary" >{label} </Typography>
    <Typography component="span" display="inline" variant="body2">{value}</Typography>
  </div>
)

const NoSelectedChannel = () => renderTextItem('No channel selected')

const SelectedChannelInfo = ({ channel }) => (
  <React.Fragment>
    {renderSubtitle1('Data')}
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
      {renderSubtitle1('Stats')}
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

const SelectedChannelDetails = ({ channel, classes }) => (
  <React.Fragment>
    <Divider className={classes.channelInfoDivider} />
    <SelectedChannelInfo channel={channel} />
    {/* TODO add component with messages */}
    <Divider className={classes.channelInfoDivider} />
    <SelectedChannelStats channel={channel} />
  </React.Fragment>
)

const SelectedChannel = ({ channel }) => {
  const classes = makeUseStylesHook()()

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
