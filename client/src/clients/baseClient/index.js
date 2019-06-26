import axios from 'axios'

const baseClient = axios.create({
  baseURL: '/api/v1',
})

export default baseClient
