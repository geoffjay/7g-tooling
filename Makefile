PROJECT := "7g-cli"

M = $(shell printf "\033[34;1m▶\033[0m")
TAG := $(shell git describe --all | sed -e's/.*\///g')

all: build

build: ; $(info $(M) Building $(PROJECT)...)
	@go build -o target/7g

build-static: ; $(info $(M) Building static binary...)
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
		-a -tags netgo -ldflags '-w -extldflags "-static"' \
		-o target/7g-static

clean: ; $(info $(M) Removing generated files... )
	@rm -rf target/

.PHONY: all build build-static clean
