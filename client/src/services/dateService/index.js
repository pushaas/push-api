import moment from 'moment'

const formatDate = (str) => moment(str).format('DD/MM/YY HH:mm:ss')

export default {
  formatDate,
}
