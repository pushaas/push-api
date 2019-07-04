import baseClient from 'clients/baseClient'

const getConfig = () => baseClient.get('/config')

export default {
  getConfig,
}
