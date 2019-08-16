VERSION=$(patsubst "%",%,$(lastword $(shell grep 'const Version' main.go)))
PRODUCT_NAME=qiic
GOVERSION=$(shell go version)
GOOS=$(word 1,$(subst /, ,$(word $(words $(GOVERSION)), $(GOVERSION))))
GOARCH=$(word 2,$(subst /, ,$(word $(words $(GOVERSION)), $(GOVERSION))))
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
BINARY_DIR=build
RELEASE_DIR=$(CURDIR)/release/$(VERSION)
BINARY_NAME=qiic

.PHONY: build test clean

all: build test

build:
	$(GOBUILD) -o $(BINARY_DIR)/$(PRODUCT_NAME)_$(GOOS)_$(GOARCH)/$(BINARY_NAME) -v

test:
	$(GOTEST) -v ./...

clean:
	$(GOCLEAN)
	rm -rf $(BINARY_DIR)

docker-build:
	docker build -t $(PRODUCT_NAME):$(VERSION) .

docker-run:
	docker build -t $(PRODUCT_NAME):$(VERSION) . \
		&& docker run --rm \
		$(PRODUCT_NAME):$(VERSION)

$(RELEASE_DIR):
	@mkdir -p $@

release-darwin-amd64:
	@$(MAKE) build release-zip GOOS=darwin GOARCH=amd64

release-zip: $(RELEASE_DIR)
	@echo " * Creating zip for $(GOOS)/$(GOARCH)"
	cd $(BINARY_DIR) && zip -9 $(RELEASE_DIR)/$(PRODUCT_NAME)_$(GOOS)_$(GOARCH).zip $(PRODUCT_NAME)_$(GOOS)_$(GOARCH)/*

release-files: release-darwin-amd64
