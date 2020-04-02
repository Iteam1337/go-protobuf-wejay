package types

import (
	"github.com/Iteam1337/go-protobuf-wejay/message"
	"github.com/Iteam1337/go-protobuf-wejay/version"
	"github.com/golang/protobuf/proto"
)

// MessageType …
type MessageType interface {
	Inv() MessageType
	Message() proto.Message
	String() string
	ByteAndVersion() [2]byte
}

// InputType …
type InputType byte

// InputTypeEnum …
const (
	ICallbackURL   InputType = 'a'
	IDeleteUser    InputType = 'b'
	IUserExists    InputType = 'c'
	INewUser       InputType = 'd'
	IPing          InputType = 'e'
	IJoinRoom      InputType = 'f'
	ILeaveRoom     InputType = 'g'
	IQueryRooms    InputType = 'h'
	IRoomExists    InputType = 'i'
	IUserLeaveRoom InputType = 'j'
	IUserRoom      InputType = 'k'
)

// Inv Response type
func (i InputType) Inv() (r MessageType) {
	switch i {
	case ICallbackURL:
		r = RCallbackURL
	case IDeleteUser:
		r = RDeleteUser
	case IUserExists:
		r = RUserExists
	case INewUser:
		r = RNewUser
	case IPing:
		r = RPong
	case IJoinRoom:
		r = RJoinRoom
	case ILeaveRoom:
		r = RLeaveRoom
	case IQueryRooms:
		r = RQueryRooms
	case IRoomExists:
		r = RRoomExists
	case IUserLeaveRoom:
		r = RUserLeaveRoom
	case IUserRoom:
		r = RUserRoom
	}
	return
}

// Message protobuf message
func (i InputType) Message() (pb proto.Message) {
	switch i {
	case ICallbackURL:
		pb = &message.CallbackURL{}
	case IDeleteUser:
		pb = &message.DeleteUser{}
	case IUserExists:
		pb = &message.UserExists{}
	case INewUser:
		pb = &message.NewUser{}
	case IPing:
		pb = &message.Ping{}
	case IJoinRoom:
		pb = &message.JoinRoom{}
	case ILeaveRoom:
		pb = &message.LeaveRoom{}
	case IQueryRooms:
		pb = &message.QueryRooms{}
	case IRoomExists:
		pb = &message.RoomExists{}
	case IUserLeaveRoom:
		pb = &message.UserLeaveRoom{}
	case IUserRoom:
		pb = &message.UserRoom{}
	}
	return
}

func (i InputType) String() string {
	return string(i)
}

// ByteAndVersion …
func (i InputType) ByteAndVersion() [2]byte {
	return [2]byte{byte(i), byte(version.Version)}
}

// ResponseType …
type ResponseType byte

// ResponseTypeEnum …
const (
	RCallbackURL   ResponseType = 'A'
	RDeleteUser    ResponseType = 'B'
	RUserExists    ResponseType = 'C'
	RNewUser       ResponseType = 'D'
	RPong          ResponseType = 'E'
	RJoinRoom      ResponseType = 'F'
	RLeaveRoom     ResponseType = 'G'
	RQueryRooms    ResponseType = 'H'
	RRoomExists    ResponseType = 'I'
	RUserLeaveRoom ResponseType = 'J'
	RUserRoom      ResponseType = 'K'
)

// Inv Input type
func (r ResponseType) Inv() (i MessageType) {
	switch r {
	case RCallbackURL:
		i = ICallbackURL
	case RDeleteUser:
		i = IDeleteUser
	case RUserExists:
		i = IUserExists
	case RNewUser:
		i = INewUser
	case RPong:
		i = IPing
	case RJoinRoom:
		i = IJoinRoom
	case RLeaveRoom:
		i = ILeaveRoom
	case RQueryRooms:
		i = IQueryRooms
	case RRoomExists:
		i = IRoomExists
	case RUserLeaveRoom:
		i = IUserLeaveRoom
	case RUserRoom:
		i = IUserRoom
	}
	return
}

// Message protobuf message
func (r ResponseType) Message() (pb proto.Message) {
	switch r {
	case RCallbackURL:
		pb = &message.CallbackURLResponse{}
	case RDeleteUser:
		pb = &message.DeleteUserResponse{}
	case RUserExists:
		pb = &message.UserExistsResponse{}
	case RNewUser:
		pb = &message.NewUserResponse{}
	case RPong:
		pb = &message.Pong{}
	case RJoinRoom:
		pb = &message.JoinRoomResponse{}
	case RLeaveRoom:
		pb = &message.LeaveRoomResponse{}
	case RQueryRooms:
		pb = &message.QueryRoomsResponse{}
	case RRoomExists:
		pb = &message.RoomExistsResponse{}
	case RUserLeaveRoom:
		pb = &message.UserLeaveRoomResponse{}
	case RUserRoom:
		pb = &message.UserRoomResponse{}
	}
	return
}

func (r ResponseType) String() string {
	return string(r)
}

// ByteAndVersion …
func (r ResponseType) ByteAndVersion() [2]byte {
	return [2]byte{byte(r), byte(version.Version)}
}
