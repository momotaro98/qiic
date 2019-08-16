VERSION=$(patsubst "%",%,$(lastword $(shell grep 'const Version' main.go)))
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
BINARY_DIR=build
BINARY_NAME=qiic

.PHONY: build test clean

all: build test

build:
	$(GOBUILD) -o $(BINARY_DIR)/$(BINARY_NAME) -v

test:
	$(GOTEST) -v ./...

clean:
	$(GOCLEAN)
	rm -rf $(BINARY_DIR)

docker-build:
	docker build -t qiic:$(VERSION) .

docker-run:
	docker build -t qiic:$(VERSION) . \
		&& docker run --rm \
		qiic:$(VERSION)
