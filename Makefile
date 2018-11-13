.PHONY: all test clean build

VERBOSE_1 := -v
VERBOSE_2 := -v -x

all: build test

build: 
	$(BUILD_ENV_FLAGS) go build $(VERBOSE_$(V)) -o bin/docker-tail .

test:     
	go test -v ./...

clean: 
	go clean ./...
	rm -rf bin

help:
	@echo "Influential make variables"
	@echo "  V                 - Build verbosity {0,1,2}."
	@echo "  BUILD_ENV_FLAGS   - Environment added to 'go build'."