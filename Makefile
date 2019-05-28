VERSION := `cat VERSION`
PKG = github.com/mradile/rssfeeder

all: vet lint cover build

.PHONY: build
build:
	CGO_ENABLED=0 go build -i -v -o release/rssfeeder/rssfeeder -ldflags="-X main.version=${VERSION}" cmd/client/*.go
	CGO_ENABLED=0 go build -i -v -o release/rssfeeder/rssfeederd -ldflags="-X main.version=${VERSION}" cmd/server/*.go

.PHONY: vet
vet:
	go vet cmd/server/*.go && go vet cmd/client/*.go

.PHONY: lint
lint:
	golint ./...

.PHONY: test
test:
	go test ./...

.PHONY: cover
cover:
	go test -coverprofile=cover.out ./...
	go tool cover -func=cover.out

.PHONY: clean
clean:
	rm -rf release/*
	go clean -testcache

