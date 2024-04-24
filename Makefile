NAME     := palworld-manager
VERSION  := $(shell cat VERSION)
REVISION := $(shell git rev-parse --short HEAD)

SRCS    := $(shell find . -type f -name '*.go')
LDFLAGS := "-s -w -X main.version=${VERSION} -X main.revision=${REVISION}"
GOOS    := $(shell go env GOOS)
GOBIN   := $(shell go env GOPATH)/bin

bin/$(NAME): $(SRCS)
	cd cmd/$(NAME); go build -ldflags=$(LDFLAGS) -o ../../bin/$(NAME) .

.PHONY: cross
cross: $(GOBIN)/goxz $(SRCS)
	cd cmd/$(NAME); goxz -n $(NAME) -pv=v$(VERSION) -build-ldflags=$(LDFLAGS) -d ../../goxz .

.PHONY: clean
clean:
	rm -rf bin/*
	rm -rf goxz/*
	rm -rf vendor/*

