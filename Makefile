MAKEFLAGS += --silent

GOPATH?=$(HOME)/go
CURDIR=$(shell pwd)
BUILD_PARAMS:=CGO_ENABLED=0 CC=gcc

ifeq ($(findstring $(GOPATH),$(CURDIR)),$(GOPATH))
    export GO111MODULE=on
endif

.PHONY: build
build:
	if [ -d "./bin/" ]; then \
		$(BUILD_PARAMS) go build -o ./bin/gmail-checker ./cmd/...; \
	else \
		mkdir -p ./bin/; \
		$(BUILD_PARAMS) go build -o ./bin/gmail-checker ./cmd/...; \
	fi
	

.PHONY: dep
dep:
	go mod tidy && go mod vendor

# TODO: make install, init