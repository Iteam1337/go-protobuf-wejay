package main

import (
	"fmt"
	"testing"

	"github.com/Iteam1337/go-protobuf-wejay/types"
)

func errInputType(name string, expected byte, i types.InputType) string {
	return fmt.Sprintf("%s was expected to be '%s', got: '%s", name, string(expected), string(i))
}

func TestInputTypes(t *testing.T) {
	if types.IPing != 'p' {
		t.Error(errInputType("IPing", 'p', types.IPing))
	}
	if types.IAction != 'a' {
		t.Error(errInputType("IAction", 'a', types.IAction))
	}
	if types.INewUser != 'n' {
		t.Error(errInputType("INewUser", 'n', types.INewUser))
	}
	if types.IUserExists != 'e' {
		t.Error(errInputType("IUserExists", 'e', types.IUserExists))
	}
	if types.ICallbackURL != 'c' {
		t.Error(errInputType("ICallbackURL", 'c', types.ICallbackURL))
	}
}

func errResponseType(name string, expected byte, r types.ResponseType) string {
	return fmt.Sprintf("%s was expected to be '%s', got: '%s", name, string(expected), string(r))
}

func TestResponseType(t *testing.T) {
	if types.RPong != 'P' {
		t.Error(errResponseType("RPong", 'P', types.RPong))
	}
	if types.RAction != 'A' {
		t.Error(errResponseType("RAction", 'A', types.RAction))
	}
	if types.RNewUser != 'N' {
		t.Error(errResponseType("RNewUser", 'N', types.RNewUser))
	}
	if types.RUserExists != 'E' {
		t.Error(errResponseType("RUserExists", 'E', types.RUserExists))
	}
	if types.RCallbackURL != 'C' {
		t.Error(errResponseType("RCallbackURL", 'C', types.RCallbackURL))
	}
}

func errInverse(name string, i types.InputType, r types.ResponseType) string {
	return fmt.Sprintf("%s in reverse was expected to be '%s', got: '%s'", name, string(i.Inv()), string(r))
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
