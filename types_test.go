package main

import (
	"fmt"
	"testing"

	"github.com/Iteam1337/go-protobuf-wejay/types"
)

type testTypes struct {
	b byte
	n string
	i types.MessageType
}

func errType(name string, expected byte, i types.MessageType) string {
	return fmt.Sprintf("%s was expected to be '%s', got: '%s", name, string(expected), i.String())
}

func TestInputTypes(t *testing.T) {
	for _, x := range []testTypes{
		{'a', "ICallbackURL", types.ICallbackURL},
		{'b', "IDeleteUser", types.IDeleteUser},
		{'c', "IUserExists", types.IUserExists},
		{'d', "INewUser", types.INewUser},
		{'e', "IPing", types.IPing},
		{'f', "IJoinRoom", types.IJoinRoom},
		{'g', "ILeaveRoom", types.ILeaveRoom},
		{'h', "IQueryRooms", types.IQueryRooms},
		{'i', "IRoomExists", types.IRoomExists},
	} {
		if x.i.String() != string(x.b) {
			t.Error(errType(x.n, x.b, x.i))
		}
	}
}

func TestResponseType(t *testing.T) {
	for _, x := range []testTypes{
		{'A', "RCallbackURL", types.RCallbackURL},
		{'B', "RDeleteUser", types.RDeleteUser},
		{'C', "RUserExists", types.RUserExists},
		{'D', "RNewUser", types.RNewUser},
		{'E', "RPong", types.RPong},
		{'F', "RJoinRoom", types.RJoinRoom},
		{'G', "RLeaveRoom", types.RLeaveRoom},
		{'H', "RQueryRooms", types.RQueryRooms},
		{'I', "RRoomExists", types.RRoomExists},
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
		{"CallbackURL", types.ICallbackURL, types.RCallbackURL},
		{"DeleteUser", types.IDeleteUser, types.RDeleteUser},
		{"UserExists", types.IUserExists, types.RUserExists},
		{"NewUser", types.INewUser, types.RNewUser},
		{"Ping", types.IPing, types.RPong},
		{"JoinRoom", types.IJoinRoom, types.RJoinRoom},
		{"LeaveRoom", types.ILeaveRoom, types.RLeaveRoom},
		{"QueryRooms", types.IQueryRooms, types.RQueryRooms},
		{"RoomExists", types.IRoomExists, types.RRoomExists},
	} {
		if x.a.Inv() != x.b {
			t.Error(errInverse(x.n, x.a.Inv(), x.b))
		}
	}
}
