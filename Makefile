TAG := latest
CONTAINER := push-api
IMAGE := rafaeleyng/$(CONTAINER)
IMAGE_TAGGED := $(IMAGE):$(TAG)
NETWORK := push-service-network
PORT_CONTAINER := 8080
PORT_HOST := 8080

CONTAINER_DEV := $(CONTAINER)-dev
IMAGE_DEV := rafaeleyng/$(CONTAINER_DEV)
IMAGE_TAGGED_DEV := $(IMAGE_DEV):$(TAG)

########################################
# app
########################################
.PHONY: setup
setup:
	@go get github.com/oxequa/realize

.PHONY: clean
clean:
	@rm -fr ./dist

.PHONY: build
build: clean
	@go build -o ./dist/push-api main.go

.PHONY: run
run:
	@go run main.go

.PHONY: kill
kill:
	@-killall push-api

.PHONY: watch
watch: kill
	@realize start

.PHONY: build-client
build-client:
	@cd client && yarn build

########################################
# docker
########################################

# dev
.PHONY: docker-clean-dev
docker-clean-dev:
	@-docker rm -f $(CONTAINER_DEV)

.PHONY: docker-build-dev
docker-build-dev:
	@docker build \
		-f Dockerfile-dev \
		-t $(IMAGE_TAGGED_DEV) \
		.

.PHONY: docker-run-dev
docker-run-dev: docker-clean-dev
	@docker run \
		-it \
		--name=$(CONTAINER_DEV) \
		--network=$(NETWORK) \
		-p $(PORT_HOST):$(PORT_CONTAINER) \
		$(IMAGE_TAGGED_DEV)

.PHONY: docker-build-and-run-dev
docker-build-and-run-dev: docker-build-dev docker-run-dev

# prod
.PHONY: docker-clean-prod
docker-clean-prod:
	@-docker rm -f $(CONTAINER)

.PHONY: docker-build-prod
docker-build-prod:
	@docker build \
		-f Dockerfile-prod \
		-t $(IMAGE_TAGGED) \
		.

.PHONY: docker-run-prod
docker-run-prod: docker-clean-prod
	@docker run \
		-e PUSHAPI_REDIS__URL="redis://push-redis:6379" \
		-it \
		--name=$(CONTAINER) \
		--network=$(NETWORK) \
		-p $(PORT_HOST):$(PORT_CONTAINER) \
		$(IMAGE_TAGGED)

.PHONY: docker-build-and-run-prod
docker-build-and-run-prod: docker-build-prod docker-run-prod

.PHONY: docker-push-prod
docker-push-prod: docker-build-prod
	@docker push \
		$(IMAGE_TAGGED)
