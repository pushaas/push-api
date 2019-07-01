import React from 'react'

export const UserContext = React.createContext(null)

export const withUser = (Component) => {
  const Wrapped = ({ value }) => (
    <UserContext.Consumer>
      {value => <Component user={value} />}
    </UserContext.Consumer>
  )
  return Wrapped
}
