VERSION := `cat VERSION`
PKG = github.com/mradile/rssfeeder

all: vet lint cover build

.PHONY: build
build:
	CGO_ENABLED=0 go build -i -v -o release/rssfeeder -ldflags="-X main.version=${VERSION}" cmd/rssfeeder/*.go
	CGO_ENABLED=0 go build -i -v -o release/rssfeederd -ldflags="-X main.version=${VERSION}" cmd/rssfeederd/*.go

.PHONY: install
install:
	cd cmd/rssfeeder && go install
	cd cmd/rssfeederd && go install

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

.PHONY: generate
generate:
	rm -f pkg/server/mock/*
	go generate ./...
	#go generate cmd/rssfeederd/main.go

.PHONY: clean
clean:
	rm -rf release/*
	go clean -testcache

.PHONY: docker
docker:
	docker build -t mradile/rssfeeder .

.PHONY: docker-publish
docker-publish: docker
	docker tag mradile/rssfeeder mradile/rssfeeder:latest
	docker tag mradile/rssfeeder mradile/rssfeeder:${VERSION}
	docker push mradile/rssfeeder:latest
	docker push mradile/rssfeeder:${VERSION}
