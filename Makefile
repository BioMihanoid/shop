.PHONY: run test lint build check

run:
	go run ./cmd/shop

build:
	go build -o bin/shop ./cmd/shop

lint:
	golangci-lint run

test:
	go test ./...

blt: build lint test