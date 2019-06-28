import React from 'react'

const SelectedChannel = (channel) => channel ? (
  <p>Select a channel</p>
) : (
  <p>{channel}</p>
)

export default SelectedChannel
