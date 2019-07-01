import axios from 'axios'

import credentialsService from 'services/credentialsService'

const baseClient = axios.create({
  baseURL: '/api/v1',
})

baseClient.interceptors.request.use((config) => {
    const credentials = credentialsService.getCredentials()
    return {
      ...config,
      auth: config.auth || credentials,
    }
  }, (error) => Promise.reject(error))

export default baseClient
