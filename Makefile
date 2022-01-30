SOURCES = $(shell find . -name '*.go' | grep -v '_test.go$$' )
APPLICATION ?= webup
VERSION ?= 

GOOS ?= linux
GOARCH ?= amd64

.DEFAULT_GOAL := build
.PHONY: fmt lint test test-race test-coverage sonar-qube build run-local image-build image-push

# Capture output and force failure when there is non-empty output
fmt:
	@echo gofmt -l .
	@OUTPUT=`gofmt -l . | grep -v ^vendor/ 2>&1`; \
	if [ "$$OUTPUT" ]; then \
		echo "gofmt must be run on the following files:"; \
		echo "$$OUTPUT"; \
		exit 1; \
	fi

lint:
	@echo golint ./...
	@OUTPUT=`golint ./... 2>&1 | grep -v ^vendor/`; \
	if [ "$$OUTPUT" ]; then \
		echo "golint errors:"; \
		echo "$$OUTPUT"; \
		exit 1; \
	fi

test:
	go test -timeout 45s ./...

test-race:
	go test -race -timeout 60s -v ./...

test-ci:
	act -l
	act -n

build: fmt lint
	rm -f bin/${APPLICATION} && env GOOS=${GOOS} GOARCH=${GOARCH} go build -o bin/${APPLICATION}
