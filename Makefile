IMAGE 		?= baltuonis/prometheus-bigbluebutton-exporter
VERSION 	= $(shell git describe --always --tags --dirty)
GO_PACKAGES = $(shell go list ./... | grep -v /vendor/)

all: format build test

# TODO: tests
# test:
# 	@echo ">> running tests"
# 	@go test $(GO_PACKAGES)

format:
	@echo ">> formatting code"
	@go fmt $(GO_PACKAGES)

build:
	@echo ">> building binaries"
	@go build

docker:
	@echo ">> building docker image"
	@docker build \
		--build-arg SOURCE_COMMIT="$(VERSION)" \
		-t $(IMAGE):$(VERSION) \
		.
	docker tag $(IMAGE):$(VERSION) $(IMAGE):latest
	docker tag $(IMAGE):$(VERSION) $(IMAGE):0.2

version:
	echo $(DOCKER_IMAGE_TAG)

.PHONY: all format build test docker
