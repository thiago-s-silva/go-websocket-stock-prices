.PHONY: default run build test docs clean

# Variables
APP_NAME=go-websocket-stock-prices

# Tasks
default: run

run:
	@go run cmd/main.go
build:
	@go build -o $(APP_NAME) cmd/main.go
test:
	@go test ./ ...