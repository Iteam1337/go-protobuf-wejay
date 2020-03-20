.DEFAULT_GOAL := main

.PHONY: all main clean go action message

clean:
	@rm -f main/*

main:
	@mkdir -p main
	@cd src; protoc --go_out=paths=source_relative:../main *.proto
	@echo "generated all"

action:
	@mkdir -p main
	@cd src; protoc --go_out=paths=source_relative:../main action.proto
	@echo "generated: action"

src/action.proto: action
main/action.pb.go: action

message:
	@mkdir -p main
	@cd src; protoc --go_out=paths=source_relative:../main message.proto
	@echo "generated: message"

src/message.proto: message
main/message.pb.go: message
