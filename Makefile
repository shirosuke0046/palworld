NAME     := palworld-manager
VERSION  := $(shell cat VERSION)
REVISION := $(shell git rev-parse --short HEAD)

SRCS    := $(shell find . -type f -name '*.go')
LDFLAGS := "-s -w -X main.version=${VERSION} -X main.revision=${REVISION}"
GOOS    := $(shell go env GOOS)

bin/$(NAME): $(SRCS)
	cd cmd/$(NAME); go build -ldflags=$(LDFLAGS) -o ../../bin/$(NAME)

.PHONY: clean
clean:
	rm -rf bin/*
	rm -rf vendor/*
