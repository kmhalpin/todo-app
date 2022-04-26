include .env.example
-include .env
export

BIN=server

.PHONY: all
all: init test build

.PHONY: init
init:
	@go mod tidy -v
	@go get -v ./...

.PHONY: test
test:
	@go test -cover -covermode atomic -coverprofile cover.out ./...

.PHONY: build
build:
	@go build -o bin/${BIN} ./cmd/server

.PHONY: run
run:
	@bin/${BIN}

.PHONY: gen
gen:
	@go generate -n -x ./...

.PHONY: env
env:
	@cat .env.example > .env
