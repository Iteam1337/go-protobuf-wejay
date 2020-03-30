.DEFAULT_GOAL := all

.PHONY: all module clean test tests message src

clean:
	@rm -f message/*

all: test module

module:
	@mkdir -p message
	@cd src; protoc --go_out=paths=source_relative:../message *.proto
	@echo "generated all"

message: module
src: module

test:
	@go test ./

tests: test
