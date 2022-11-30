.DEFAULT_GOAL := build

SHELL  = /bin/bash
export GOBIN := $(PWD)/bin
export PATH  := $(GOBIN):$(PATH)

install-tools:
	awk '/\s_ ".+"/ {print $$2}' $(PWD)/tools/tools.go | xargs go install
.PHONY:install-tools

fmt:
	go fmt ./...
.PHONY:fmt

lint: fmt
	${GOBIN}/golangci-lint run ./...
.PHONY:lint

vet: fmt
	go vet ./...
.PHONY:vet

build: vet
	go build -o bin/climer .
.PHONY:build
