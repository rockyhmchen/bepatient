VERSION=0.1.0
BINARY_NAME=bepatient

GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOGET=$(GOCMD) get
GOTEST=gotestsum

all: clean tidy test build

build: 
	$(GOBUILD) -o $(BINARY_NAME) -v -ldflags="-X 'github.com/rockyhmchen/bepatient/version.Version=$(VERSION)'"

build_linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_NAME) -v -ldflags="-X 'github.com/rockyhmchen/bepatient/version.Version=$(VERSION)'"

test: 
	$(GOGET) gotest.tools/gotestsum
	$(GOTEST) --format short-verbose

clean: 
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)

run:
	$(GOBUILD) -o $(BINARY_NAME) -v ./...
	./$(BINARY_NAME)

tidy:
	$(GOCMD) mod tidy