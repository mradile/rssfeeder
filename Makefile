PKG = github.com/mradile/rssfeeder

all: vet lint cover build

.PHONY: build
build:
	CGO_ENABLED=0 go build -i -v -o release/rssfeeder cmd/rssfeeder/*.go
	CGO_ENABLED=0 go build -i -v -o release/rssfeederd cmd/rssfeederd/*.go

.PHONY: install
install:
	cd cmd/rssfeeder && go install
	cd cmd/rssfeederd && go install

.PHONY: vet
vet:
	go vet cmd/rssfeeder/*.go && go vet cmd/rssfeederd/*.go

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

