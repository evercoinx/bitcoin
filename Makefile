.PHONY: test

build:
	GOFLAGS=-mod=vendor go build -o bin/cli cmd/cli/main.go

run: build
	./bin/cli

test:
	go test -count=1 ./...
