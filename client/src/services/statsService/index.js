import baseClient from 'clients/baseClient'

const getGlobalStats = () => baseClient.get('/stats/global')
  .then(({ data }) => data)

export default {
  getGlobalStats,
}
