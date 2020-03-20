.DEFAULT_GOAL := module

.PHONY: all module clean go action newUser callbackURL

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

newUser:
	@mkdir -p message
	@cd src; protoc --go_out=paths=source_relative:../message newUser.proto
	@echo "generated: newUser"

src/newUser.proto: newUser
message/newUser.pb.go: newUser

userExists:
	@mkdir -p message
	@cd src; protoc --go_out=paths=source_relative:../message userExists.proto
	@echo "generated: userExists"

src/userExists.proto: userExists
message/userExists.pb.go: userExists


callbackURL:
	@mkdir -p message
	@cd src; protoc --go_out=paths=source_relative:../message callbackURL.proto
	@echo "generated: callbackURL"

src/callbackURL.proto: callbackURL
message/callbackURL.pb.go: callbackURL
