VERSION := $(shell git describe --tags --abbrev=0)
BUILDTIME := $(shell date -u '+%Y-%m-%dT%H:%M:%SZ')

GOLDFLAGS += -X main.Version=$(VERSION)
GOLDFLAGS += -X main.Buildtime=$(BUILDTIME)
GOFLAGS = -ldflags "$(GOLDFLAGS)"

run: build
	./gpf

dry-run: build
	./gpf -dry-run

build:
	go build -o gpf $(GOFLAGS) .

test: build
	go test -v ./...
