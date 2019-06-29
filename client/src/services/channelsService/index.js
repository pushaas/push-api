import baseClient from 'clients/baseClient'

const getChannels = () => baseClient.get('/channels')
  .then(({ data }) => data)

const deleteChannel = (id) => baseClient.delete(`/channels/${id}`)
  .then(({ data }) => data)

export default {
  getChannels,
  deleteChannel,
}
