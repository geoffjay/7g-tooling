PROJECT := "7g-tooling"

M = $(shell printf "\033[34;1m▶\033[0m")
TAG := $(shell git describe --all | sed -e's/.*\///g')

all: build

docs: ; $(info $(M) Building Swagger documentation...)
	@swag init -g internal/service/main.go --parseInternal

setup: ; $(info $(M) Fetching golang build dependencies...)
	@go get -u github.com/go-bindata/go-bindata/...

bindata: ; $(info $(M) Generating binary data to package...)
	@go generate ./data/gql

build: bindata; $(info $(M) Building $(PROJECT)...)
	@go build -o target/7g

build-static: ; $(info $(M) Building static binary...)
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
		-a -tags netgo -ldflags '-w -extldflags "-static"' \
		-o target/7g-static

clean: ; $(info $(M) Removing generated files... )
	@rm -rf target/

.PHONY: all build build-static clean docs
