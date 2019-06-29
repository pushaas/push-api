import React from 'react'
import Divider from '@material-ui/core/Divider'
import Typography from '@material-ui/core/Typography'

import dateService from 'services/dateService'

import { makeUseStylesHook } from 'components/App/styles'
import Title from 'components/common/Title'

const renderText1 = (text) => (<Typography component="p" variant="subtitle1" color="primary" gutterBottom>{text}</Typography>)
const renderText2 = (text) => (<Typography component="span" variant="subtitle2" gutterBottom>{text}</Typography>)

const renderChannel = (channel, classes) => {
  if (!channel) {
    return renderText2('No channel selected')
  }

  return (
    <React.Fragment>
      {renderText1('Data')}
      {renderText2(`id: ${channel.id}`)}
      {renderText2(`ttl: ${channel.ttl}`)}
      {renderText2(`created: ${dateService.formatDate(channel.created)}`)}

      <Divider className={classes.channelInfoDivider} />

      {renderText1('Stats')}
      {renderText2(`id: ${channel.id}`)}

      <Divider className={classes.channelInfoDivider} />

      {renderText1('Messages')}
    </React.Fragment>
  )
}

const SelectedChannel = ({ selectedChannel }) => {
  const classes = makeUseStylesHook()()

  return (
    <React.Fragment>
      <Title>
        Selected Channel
      </Title>
      {renderChannel(selectedChannel, classes)}
    </React.Fragment>
  )
}

export default SelectedChannel
