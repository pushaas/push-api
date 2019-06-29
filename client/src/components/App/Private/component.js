import React from 'react'
import { makeUseStylesHook } from 'components/App/styles'

import Header from './Header'
import Menu from './Menu'
import Main from './Main'

const Private = () => {
  const classes = makeUseStylesHook()()
  const [open, setOpen] = React.useState(true)
  const handleDrawerOpen = () => { setOpen(true) }
  const handleDrawerClose = () => { setOpen(false) }

  return (
    <div className={classes.root}>
      <Header open={open} handleDrawerOpen={handleDrawerOpen} />
      <Menu open={open} handleDrawerClose={handleDrawerClose} />
      <Main />
    </div>
  )
}

export default  Private
