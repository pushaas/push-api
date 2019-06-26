import baseClient from 'clients/baseClient'

const getChannels = () => baseClient.get('/channels')
  .then(({ data }) => data)

export default {
  getChannels,
}
