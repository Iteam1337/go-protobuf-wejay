package main

import (
	"fmt"
	"testing"

	"github.com/Iteam1337/go-protobuf-wejay/types"
)

type testTypes struct {
	n string
	b byte
	i types.MessageType
}

func errType(name string, expected byte, i types.MessageType) string {
	return fmt.Sprintf("%s was expected to be '%s', got: '%s", name, string(expected), i.String())
}

func TestInputTypes(t *testing.T) {
	for _, x := range []testTypes{
		{"IPing", 'p', types.IPing},
		{"IAction", 'a', types.IAction},
		{"INewUser", 'n', types.INewUser},
		{"IUserExists", 'e', types.IUserExists},
		{"ICallbackURL", 'c', types.ICallbackURL},
		{"INowPlaying", 'x', types.INowPlaying},
		{"IListen", 'l', types.IListen},
	} {
		if x.i.String() != string(x.b) {
			t.Error(errType(x.n, x.b, x.i))
		}
	}
}

func TestResponseType(t *testing.T) {
	for _, x := range []testTypes{
		{"RPong", 'P', types.RPong},
		{"RAction", 'A', types.RAction},
		{"RNewUser", 'N', types.RNewUser},
		{"RUserExists", 'E', types.RUserExists},
		{"RCallbackURL", 'C', types.RCallbackURL},
		{"RNowPlaying", 'X', types.RNowPlaying},
		{"RListen", 'L', types.RListen},
	} {
		if x.i.String() != string(x.b) {
			t.Error(errType(x.n, x.b, x.i))
		}
	}
}

type testInverse struct {
	n string
	a types.MessageType
	b types.MessageType
}

func errInverse(name string, a types.MessageType, b types.MessageType) string {
	return fmt.Sprintf("%s in reverse was expected to be '%s', got: '%s'", name, a.Inv().String(), b.String())
}

func TestInverse(t *testing.T) {
	for _, x := range []testInverse{
		{"IPing", types.IPing, types.RPong},
		{"IAction", types.IAction, types.RAction},
		{"INewUser", types.INewUser, types.RNewUser},
		{"IUserExists", types.IUserExists, types.RUserExists},
		{"ICallbackURL", types.ICallbackURL, types.RCallbackURL},
		{"INowPlaying", types.INowPlaying, types.RNowPlaying},
	} {
		if x.a.Inv() != x.b {
			t.Error(errInverse(x.n, x.a, x.b))
		}
	}
}
