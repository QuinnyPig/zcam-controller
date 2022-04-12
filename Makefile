.PHONY: all
all: build

.PHONY: build
build:
	go build -v ./...

.PHONY: test
test:
	go test -coverprofile=coverage.out -timeout=20s -v ./...
	go tool cover -func=coverage.out
