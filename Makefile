VERSION := `cat VERSION`
PKG = github.com/mradile/rssfeeder

.PHONY: build
build:
	CGO_ENABLED=0 go build -i -v -o release/rssfeeder/rssfeeder -ldflags="-X main.version=${VERSION}" cmd/client/*.go
	CGO_ENABLED=0 go build -i -v -o release/rssfeeder/rssfeederd -ldflags="-X main.version=${VERSION}" cmd/server/*.go

.PHONY: clean
clean:
	rm -rf release/*
	go clean -testcache

