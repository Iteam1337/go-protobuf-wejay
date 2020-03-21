.DEFAULT_GOAL := module

.PHONY: all module clean test message src

clean:
	@rm -f message/*

module: test
	@mkdir -p message
	@cd src; protoc --go_out=paths=source_relative:../message *.proto
	@echo "generated all"

message: module
message: src

test:
	@go test ./

