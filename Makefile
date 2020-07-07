PROJECT := "7g-tooling"

M = $(shell printf "\033[34;1m▶\033[0m")
TAG := $(shell git describe --all | sed -e's/.*\///g')
GBD := $(shell command -v go-bindata 2> /dev/null)
AIR := $(shell command -v air 2> /dev/null)

all: build

docs: ; $(info $(M) Building Swagger documentation...)
	@swag init -g internal/service/main.go --parseInternal

setup: ; $(info $(M) Fetching golang build dependencies...)
ifndef GBD
	@go get -u github.com/go-bindata/go-bindata/...
endif

bindata: setup; $(info $(M) Generating binary data to package...)
	@go generate ./data/gql

build: bindata; $(info $(M) Building $(PROJECT)...)
	@go build -o target/7g

build-static: ; $(info $(M) Building static binary...)
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
		-a -tags netgo -ldflags '-w -extldflags "-static"' \
		-o target/7g-static

server: ; $(info $(M) Run service using hot-reloading...)
ifndef AIR
	@go get github.com/cosmtrek/air
endif
	@air

# `dotenv` is installed with `pip install "python-dotenv[cli]"`
release: ; $(info $(M) Create release...)
	@dotenv -f $(USER)/.env.gh goreleaser --rm-dist

test: bindata; $(info $(M) Running tests...)
	@go test -v -tags integration ./...

clean: ; $(info $(M) Removing generated files... )
	@rm -rf target/

.PHONY: all docs build build-static server test clean
