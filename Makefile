.PHONY: build \
	run

########################################
# app
########################################
setup:
	@go get github.com/oxequa/realize

clean:
	@rm -fr ./dist

build: clean
#	@cp ./config/$(ENV).yml ./dist/config.yml
	@go build -o ./dist/push-api main.go

run:
	@PUSHAPI_ENV=local go run main.go

watch:
	@PUSHAPI_ENV=local realize start --run --no-config

########################################
# docker
########################################

# dev
docker-build-dev:
	@docker build \
		-f Dockerfile-dev \
		-t push-api:latest \
		.

docker-run-dev:
	@docker run \
		-it \
		-p 9000:9000 \
		push-api:latest

docker-build-and-run-dev: docker-build-dev docker-run-dev

# prod
docker-build-prod:
	@docker build \
		-f Dockerfile-prod \
		-t rafaeleyng/push-api:latest \
		.

docker-run-prod:
	@docker run \
		-it \
		-p 9000:9000 \
		rafaeleyng/push-api:latest

docker-build-and-run-prod: docker-build-prod docker-run-prod

docker-push-prod: docker-build-prod
	@docker push \
		rafaeleyng/push-api
