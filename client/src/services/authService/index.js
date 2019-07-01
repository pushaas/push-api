import baseClient from 'clients/baseClient'

const checkAuth = ({ username, password } = {}) => {
  const config = {
    auth: {
      username,
      password,
    },
  }

  return baseClient.get('/auth', config)
    .then(({ data }) => data)
}

export default {
  checkAuth,
}
