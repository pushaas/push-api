# push-api

This is documented [here](https://github.com/pushaas/pushaas-docs#component-push-api)

`push-api` exposes an API and a web client to work with the [`push-service`](https://github.com/pushaas/push-service).

The default credentials to access the web client are [here](https://github.com/pushaas/push-api/blob/master/push-api/ctors/config.go) (`"api.basic_auth_user"` and `"api.basic_auth_password"`).

## running locally

Requires [push-redis](https://github.com/pushaas/push-redis) to be running.

- run the API:
	```shell
	make run
	```
- run the client (in other shell):
	```shell
	cd client
	yarn
	yarn start
	```

## publishing images

```shell
make docker-push TAG=<tag>
```

---
