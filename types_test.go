package main

import (
	"fmt"
	"testing"

	"github.com/Iteam1337/go-protobuf-wejay/types"
)

func errType(name string, expected byte, i types.MessageType) string {
	return fmt.Sprintf("%s was expected to be '%s', got: '%s", name, string(expected), i.String())
}

func TestInputTypes(t *testing.T) {
	if types.IPing != 'p' {
		t.Error(errType("IPing", 'p', types.IPing))
	}
	if types.IAction != 'a' {
		t.Error(errType("IAction", 'a', types.IAction))
	}
	if types.INewUser != 'n' {
		t.Error(errType("INewUser", 'n', types.INewUser))
	}
	if types.IUserExists != 'e' {
		t.Error(errType("IUserExists", 'e', types.IUserExists))
	}
	if types.ICallbackURL != 'c' {
		t.Error(errType("ICallbackURL", 'c', types.ICallbackURL))
	}
}

func TestResponseType(t *testing.T) {
	if types.RPong != 'P' {
		t.Error(errType("RPong", 'P', types.RPong))
	}
	if types.RAction != 'A' {
		t.Error(errType("RAction", 'A', types.RAction))
	}
	if types.RNewUser != 'N' {
		t.Error(errType("RNewUser", 'N', types.RNewUser))
	}
	if types.RUserExists != 'E' {
		t.Error(errType("RUserExists", 'E', types.RUserExists))
	}
	if types.RCallbackURL != 'C' {
		t.Error(errType("RCallbackURL", 'C', types.RCallbackURL))
	}
}

func errInverse(name string, a types.MessageType, b types.MessageType) string {
	return fmt.Sprintf("%s in reverse was expected to be '%s', got: '%s'", name, a.Inv().String(), b.String())
}

func TestInverse(t *testing.T) {
	if types.IPing.Inv() != types.RPong {
		t.Error(errInverse("IPing", types.IPing, types.RPong))
	}
	if types.IAction.Inv() != types.RAction {
		t.Error(errInverse("IAction", types.IAction, types.RAction))
	}
	if types.INewUser.Inv() != types.RNewUser {
		t.Error(errInverse("INewUser", types.INewUser, types.RNewUser))
	}
	if types.IUserExists.Inv() != types.RUserExists {
		t.Error(errInverse("IUserExists", types.IUserExists, types.RUserExists))
	}
	if types.ICallbackURL.Inv() != types.RCallbackURL {
		t.Error(errInverse("ICallbackURL", types.ICallbackURL, types.RCallbackURL))
	}
}
