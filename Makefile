GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
BINARY_NAME=migraine
BINARY_PATH=cli/main.go
MIGRAINES_DIR=./migraines

all: test build
build:
				$(GOBUILD) -o $(BINARY_NAME) -v ${BINARY_PATH}
test:
				$(GOTEST) -v ./...
clean:
				$(GOCLEAN)
				rm -f $(BINARY_NAME)
				rm -rf ${MIGRAINES_DIR}

