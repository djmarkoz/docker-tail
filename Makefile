.PHONY: all test clean build

VERSION ?= snapshot
VERBOSE_1 := -v
VERBOSE_2 := -v -x

all: build test

build: 
	$(BUILD_ENV_FLAGS) go build $(VERBOSE_$(V)) -o bin/docker-tail-$(VERSION)

test:     
	go test -v ./...

clean: 
	go clean ./...
	rm -rf bin

dist:
	go get github.com/mitchellh/gox
	gox -os="linux darwin windows freebsd" -output="bin/{{.Dir}}-$(VERSION)-{{.OS}}-{{.Arch}}"

help:
	@echo "Influential make variables"
	@echo "  VERSION		   - Version used for build."
	@echo "  V                 - Build verbosity {0,1,2}."
	@echo "  BUILD_ENV_FLAGS   - Environment added to 'go build'."
