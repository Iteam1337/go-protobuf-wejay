.DEFAULT_GOAL := module

.PHONY: all module clean go action message

clean:
	@rm -f message/*

module:
	@mkdir -p message
	@cd src; protoc --go_out=paths=source_relative:../message *.proto
	@echo "generated all"

action:
	@mkdir -p message
	@cd src; protoc --go_out=paths=source_relative:../message action.proto
	@echo "generated: action"

src/action.proto: action
message/action.pb.go: action

message:
	@mkdir -p message
	@cd src; protoc --go_out=paths=source_relative:../message message.proto
	@echo "generated: message"

src/message.proto: message
message/message.pb.go: message
