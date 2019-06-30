import React from 'react'
import Table from '@material-ui/core/Table'
import TableBody from '@material-ui/core/TableBody'
import TableCell from '@material-ui/core/TableCell'
import TableHead from '@material-ui/core/TableHead'
import TableRow from '@material-ui/core/TableRow'
import IconButton from '@material-ui/core/IconButton'
import CheckIcon from '@material-ui/icons/Check'
import DeleteIcon from '@material-ui/icons/Delete'

import Title from 'components/common/Title'

import dateService from 'services/dateService'

import { useStyles } from 'components/App/styles'

const ChannelList = ({ channels, onDeleteChannel, onSelectChannel }) => {
  const classes = useStyles()

  return (
    <React.Fragment>
      <Title>
        Persistent Channels <small>({channels.length})</small>
      </Title>
      <Table size="small">
        <TableHead>
          <TableRow>
            <TableCell />
            <TableCell>Id</TableCell>
            <TableCell>TTL</TableCell>
            <TableCell>Created</TableCell>
            <TableCell />
          </TableRow>
        </TableHead>
        <TableBody>
          {channels.map(channel => (
            <TableRow key={channel.id}>
              <TableCell>
                <IconButton edge="end" size="small" onClick={() => onSelectChannel(channel)}>
                  <CheckIcon />
                </IconButton>
              </TableCell>
              <TableCell className={classes.channelIdCell}>{channel.id}</TableCell>
              <TableCell>{channel.ttl}</TableCell>
              <TableCell>{dateService.formatDate(channel.created)}</TableCell>
              <TableCell>
                <IconButton edge="end" size="small" onClick={() => onDeleteChannel(channel)}>
                  <DeleteIcon />
                </IconButton>
              </TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>
    </React.Fragment>
  )
}

export default ChannelList
