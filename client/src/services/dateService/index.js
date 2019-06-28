import moment from 'moment'

const formatDate = (str) => moment(str).format('DD/MM/YYYY HH:mm:ss')

export default {
  formatDate,
}
