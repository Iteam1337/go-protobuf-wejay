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
		{"ICallbackURL", 'c', types.ICallbackURL},
		{"IDeleteUser", 'd', types.IDeleteUser},
		{"IUserExists", 'e', types.IUserExists},
		{"INewUser", 'n', types.INewUser},
		{"IPing", 'p', types.IPing},
		{"IJoinRoom", 'j', types.IJoinRoom},
		{"ILeaveRoom", 'l', types.ILeaveRoom},
		{"IQueryRooms", 'q', types.IQueryRooms},
	} {
		if x.i.String() != string(x.b) {
			t.Error(errType(x.n, x.b, x.i))
		}
	}
}

func TestResponseType(t *testing.T) {
	for _, x := range []testTypes{
		{"RCallbackURL", 'C', types.RCallbackURL},
		{"RDeleteUser", 'D', types.RDeleteUser},
		{"RUserExists", 'E', types.RUserExists},
		{"RNewUser", 'N', types.RNewUser},
		{"RPong", 'P', types.RPong},
		{"RJoinRoom", 'J', types.RJoinRoom},
		{"RLeaveRoom", 'L', types.RLeaveRoom},
		{"RQueryRooms", 'Q', types.RQueryRooms},
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
		{"ICallbackURL", types.ICallbackURL, types.RCallbackURL},
		{"IDeleteUser", types.IDeleteUser, types.RDeleteUser},
		{"IUserExists", types.IUserExists, types.RUserExists},
		{"INewUser", types.INewUser, types.RNewUser},
		{"IPing", types.IPing, types.RPong},
		{"IJoinRoom", types.IJoinRoom, types.RJoinRoom},
		{"ILeaveRoom", types.ILeaveRoom, types.RLeaveRoom},
		{"IQueryRooms", types.IQueryRooms, types.RQueryRooms},
	} {
		if x.a.Inv() != x.b {
			t.Error(errInverse(x.n, x.a.Inv(), x.b))
		}
	}
}
