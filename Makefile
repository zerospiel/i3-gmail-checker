MAKEFLAGS += --silent

GOPATH?=$(HOME)/go
CURDIR=$(shell pwd)
BUILD=$()

ifeq ($(findstring $(GOPATH),$(CURDIR)),$(GOPATH))
    export GO111MODULE=on
endif

.PHONY: build
build:
	go build -o gmail-checker ./cmd/...

.PHONY: dep
dep:
	go mod tidy && go mod vendor