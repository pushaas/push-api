import baseClient from 'clients/baseClient'

const getChannelStats = (id) => baseClient.get(`/stats/channels/${id}`)

const getGlobalStats = () => baseClient.get('/stats/global')

export default {
  getChannelStats,
  getGlobalStats,
}
