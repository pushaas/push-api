import baseClient from 'clients/baseClient'

const getChannels = () => baseClient.get('/channels')

const deleteChannel = (id) => baseClient.delete(`/channels/${id}`)

export default {
  getChannels,
  deleteChannel,
}
