import React from 'react'
import Table from '@material-ui/core/Table'
import TableBody from '@material-ui/core/TableBody'
import TableCell from '@material-ui/core/TableCell'
import TableHead from '@material-ui/core/TableHead'
import TableRow from '@material-ui/core/TableRow'

import Title from 'components/common/Title'

import dateService from 'services/dateService'

const ChannelList = ({ channels, onSelectChannel }) => (
  <React.Fragment>
    <Title>Channels</Title>
    <Table size="small">
      <TableHead>
        <TableRow>
          <TableCell>Id</TableCell>
          <TableCell>TTL</TableCell>
          <TableCell>Created</TableCell>
        </TableRow>
      </TableHead>
      <TableBody>
        {channels.map(row => (
          <TableRow key={row.id}>
            <TableCell>{row.id}</TableCell>
            <TableCell align="right">{row.ttl}</TableCell>
            <TableCell>{dateService.formatDate(row.created)}</TableCell>
          </TableRow>
        ))}
      </TableBody>
    </Table>
  </React.Fragment>
)

export default ChannelList
