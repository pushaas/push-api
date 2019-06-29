import baseClient from 'clients/baseClient'

const getChannelStats = (id) => baseClient.get(`/stats/channels/${id}`)
  .then(({ data }) => data)

const getGlobalStats = () => baseClient.get('/stats/global')
  .then(({ data }) => data)

export default {
  getChannelStats,
  getGlobalStats,
}
