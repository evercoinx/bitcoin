.PHONY: build run test update

COVERAGE_FILE := coverage.out

build:
	GOFLAGS=-mod=vendor go build -o bin/cli cmd/cli/main.go

test:
	go test -count=1 -cover ./...

cover:
	go test -coverprofile=$(COVERAGE_FILE) -covermode=count ./...
	go tool cover -html=$(COVERAGE_FILE)

update:
	GOPRIVATE=github.com/evercoinx go get -u github.com/evercoinx/kit
	go mod vendor
