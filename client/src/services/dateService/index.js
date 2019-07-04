import moment from 'moment'

const formatDate = (str) => moment(str).format('DD/MM/YY HH:mm:ss')
const calculateExpiration = (str, ttl) => formatDate(moment(str).add(ttl, 'seconds'))

export default {
  calculateExpiration,
  formatDate,
}
