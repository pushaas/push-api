import baseClient from 'clients/baseClient'

const postMessage = (data) => baseClient.post('/messages', data)

export default {
  postMessage,
}
