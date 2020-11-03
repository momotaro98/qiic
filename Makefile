VERSION=$(patsubst "%",%,$(lastword $(shell grep 'const Version' main.go)))
PRODUCT_NAME=qiic
GOVERSION=$(shell go version)
GOOS=$(word 1,$(subst /, ,$(word $(words $(GOVERSION)), $(GOVERSION))))
GOARCH=$(word 2,$(subst /, ,$(word $(words $(GOVERSION)), $(GOVERSION))))
GOCMD=go
BINPATH=./cmd/qiic
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test

.PHONY: build test clean

all: build test

build:
	$(GOBUILD) $(BINPATH)

test:
	$(GOTEST) -v ./...

clean:
	rm -f ./qiic

docker-build:
	docker build -t $(PRODUCT_NAME):$(VERSION) .

docker-run:
	docker build -t $(PRODUCT_NAME):$(VERSION) . \
		&& docker run --rm \
		$(PRODUCT_NAME):$(VERSION)

.PHONY: help
help: ## Help command
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
