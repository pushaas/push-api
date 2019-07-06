import React, { useContext, useEffect, useState } from 'react'
import TextField from '@material-ui/core/TextField'
import url from 'url'

import dateService from 'services/dateService'
import messagesService from 'services/messagesService'
import pushStreamService from 'services/pushStreamService'

import {
  renderSubtitle1,
  renderBody2,
} from 'components/common/textHelpers'
import ConfigContext from 'components/contexts/ConfigContext'
import { useStyles } from 'components/App/Private/styles'

class Message {
  constructor({ channel, id, text, eventId, isLastMessageFromBatch, time }) {
    this.channel = channel
    this.id = id
    this.text = text
    this.eventId = eventId
    this.isLastMessageFromBatch = isLastMessageFromBatch
    this.time = time
  }
}

const Messages = ({ channel }) => {
  const classes = useStyles()
  const config = useContext(ConfigContext)
  const [message, setMessage] = useState('')
  const [messages, setMessages] = useState([])
  const [pushStreamInstance, setPushStreamInstance] = useState(null)

  const handleSetMessage = (e) => {
    setMessage(e.target.value)
  }

  const handleKeyPress = (e) => {
    if (e.key === 'Enter') {
      const apiMessage = {
        channels: [channel.id],
        content: message,
      }
      messagesService.postMessage(apiMessage)
      setMessage('')
    }
  }

  useEffect(() => {
    const parsedUrl = url.parse(config.pushStream.url)
    const settings = {
      host: parsedUrl.hostname,
      port: parsedUrl.port,
      modes: 'eventsource',
      messagesPublishedAfter: 900,
      messagesControlByArgument: true,
      onerror: (err) => console.error('[onerror]', err),
    }

    const instance = pushStreamService.newPushStreamInstance(settings)
    setPushStreamInstance(instance)
    instance.addChannel(channel.id)
    instance.connect()

    return () => {
      instance.disconnect()
    }
  }, [setPushStreamInstance, config, channel])

  useEffect(() => {
    if (!pushStreamInstance) {
      return
    }
    const onMessage = (text, id, channel, eventId, isLastMessageFromBatch, time) => {
      setMessages([...messages, new Message({ text, id, channel, eventId, isLastMessageFromBatch, time })])
    }
    pushStreamInstance.onmessage = onMessage
  }, [pushStreamInstance, messages])

  return (
    <div className={classes.code1}>
      {renderSubtitle1('Messages flowing on this channel')}
      <pre className={classes.messagesPre}>
        <code>
          {messages.map(message => (<React.Fragment key={message.id}>{`[${dateService.formatDate(message.time)}] ${message.text}\n`}</React.Fragment>))}
        </code>
      </pre>
      <TextField
        fullWidth
        value={message}
        placeholder="Type a message and press Enter"
        onChange={handleSetMessage}
        onKeyPress={handleKeyPress}
      />
      <div className={classes.messagesNote}>
        {renderBody2('Note: this message will be sent to all subscribers. If you just want to see messages flowing, just wait a minute for a "ping" message.')}
      </div>
    </div>
  )
}

export default Messages
