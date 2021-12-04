.PHONY: build run test update

build:
	GOFLAGS=-mod=vendor go build -o bin/cli cmd/cli/main.go

test:
	go test -count=1 ./...

update:
	GOPRIVATE=github.com/evercoinx go mod tidy
	go mod vendor
