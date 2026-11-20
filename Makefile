.PHONY: build test lint bench

build:
	go build -o bin/mcp-guard ./cmd/mcp-guard

test:
	go test -v -race ./...

bench:
	go test -bench=. -benchmem ./...

lint:
	golangci-lint run ./...
