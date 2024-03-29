.PHONY: all
all: build

.PHONY: build
build:
	go build

.PHONY: install
install:
	cp zcam-controller /usr/local/bin/

.PHONY: test
test:
	go test -coverprofile=coverage.out -timeout=20s -v ./...
	go tool cover -func=coverage.out
