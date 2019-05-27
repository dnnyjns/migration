GOCMD=go
GOTEST=$(GOCMD) test

all: test
test:
				$(GOTEST) -v ./...

