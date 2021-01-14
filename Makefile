default: build test

prebuild:
	@go mod tidy

build: prebuild
	@go build

.PHONY: build test prebuild
