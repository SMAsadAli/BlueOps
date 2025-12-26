GO ?= go

.PHONY: test lint build

test:
	$(GO) test ./...

lint:
	golangci-lint run

build:
	$(GO) build ./...
