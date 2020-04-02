.DEFAULT_GOAL := all

.PHONY: all module clean test tests message src

clean:
	@rm -rf message

all: test module

module:
	@mkdir -p message
	@cd src; protoc -I=. --proto_path=. --go_out=plugins=grpc,paths=source_relative:../message **.proto
	@echo "generated all"

message: module
src: module

test:
	@go test ./

tests: test
