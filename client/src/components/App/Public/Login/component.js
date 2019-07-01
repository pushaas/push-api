import React, { useContext, useState } from 'react'
import Button from '@material-ui/core/Button'
import TextField from '@material-ui/core/TextField'
import Grid from '@material-ui/core/Grid'
import Typography from '@material-ui/core/Typography'
import Container from '@material-ui/core/Container'

import { useStyles } from 'components/App/Public/styles'
import SetUserContext from 'components/contexts/SetUserContext'

const useFormInput = (initial) => {
  const [value, setValue] = useState(initial)
  const handleChange = (e) => setValue(e.target.value)
  return [value, handleChange]
}

const Login = () => {
  const classes = useStyles()
  const setUser = useContext(SetUserContext)

  const [username, setUsername] = useFormInput('')
  const [password, setPassword] = useFormInput('')

  const login = (e) => {
    e.preventDefault()
    // TODO
    console.log('### login', username, password)
    setUser({ name: 'Rafael' })
  }

  return (
    <Container component="main" maxWidth="xs">
      <div className={classes.paper}>
        <Typography component="h1" variant="h5">
          PushApi Admin
        </Typography>
        <form className={classes.form} noValidate onSubmit={login}>
          <TextField
            value={username}
            onChange={setUsername}
            variant="outlined"
            margin="normal"
            required
            fullWidth
            label="Username"
            autoComplete="email"
            autoFocus
          />
          <TextField
            value={password}
            onChange={setPassword}
            variant="outlined"
            margin="normal"
            required
            fullWidth
            label="Password"
            type="password"
            autoComplete="current-password"
          />
          <Button
            type="submit"
            fullWidth
            variant="contained"
            color="primary"
            className={classes.submit}
          >
            Sign In
          </Button>
          <Grid container>
            <Grid item xs>
              <Typography variant="body2" color="textSecondary" align="center">
                Use the same credentials your app uses to call the PushApi
              </Typography>
            </Grid>
          </Grid>
        </form>
      </div>
    </Container>
  )
}

export default Login
