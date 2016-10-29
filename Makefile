SHELL := /bin/bash

dependencies: clear & glide install
build: go build
release: dependencies build

.PHONY: build