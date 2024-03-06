.PHONY: build clean test

GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test

BINARY_NAME=sheet-watcher

all: build

build: 
	$(GOBUILD) -o $(BINARY_NAME) -v ./cmd/sheet-watcher/main.go

test: 
	$(GOTEST) -v ./...

clean: 
	$(GOCLEAN)
	rm -f $(BINARY_NAME)