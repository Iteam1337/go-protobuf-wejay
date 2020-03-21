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
	IPing        InputType = 'p'
	IAction      InputType = 'a'
	INewUser     InputType = 'n'
	IUserExists  InputType = 'e'
	ICallbackURL InputType = 'c'
	INowPlaying  InputType = 'x'
	IListen      InputType = 'l'
)

// Inv Response type
func (i InputType) Inv() (r MessageType) {
	switch i {
	case IPing:
		r = RPong
	case IAction:
		r = RAction
	case INewUser:
		r = RNewUser
	case IUserExists:
		r = RUserExists
	case ICallbackURL:
		r = RCallbackURL
	case INowPlaying:
		r = RNowPlaying
	case IListen:
		r = RListen
	}
	return
}

// Message protobuf message
func (i InputType) Message() (pb proto.Message) {
	switch i {
	case IPing:
		pb = &message.Ping{}
	case IAction:
		pb = &message.Action{}
	case INewUser:
		pb = &message.NewUser{}
	case IUserExists:
		pb = &message.UserExists{}
	case ICallbackURL:
		pb = &message.CallbackURL{}
	case INowPlaying:
		pb = &message.NowPlaying{}
	case IListen:
		pb = &message.Listen{}
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
	RPong        ResponseType = 'P'
	RAction      ResponseType = 'A'
	RNewUser     ResponseType = 'N'
	RUserExists  ResponseType = 'E'
	RCallbackURL ResponseType = 'C'
	RNowPlaying  ResponseType = 'X'
	RListen      ResponseType = 'L'
)

// Inv Input type
func (r ResponseType) Inv() (i MessageType) {
	switch r {
	case RPong:
		i = IPing
	case RAction:
		i = IAction
	case RNewUser:
		i = INewUser
	case RUserExists:
		i = IUserExists
	case RCallbackURL:
		i = ICallbackURL
	case RNowPlaying:
		i = INowPlaying
	case RListen:
		i = IListen
	}
	return
}

// Message protobuf message
func (r ResponseType) Message() (pb proto.Message) {
	switch r {
	case RPong:
		pb = &message.Pong{}
	case RAction:
		pb = &message.ActionResponse{}
	case RNewUser:
		pb = &message.NewUserResponse{}
	case RUserExists:
		pb = &message.UserExistsResponse{}
	case RCallbackURL:
		pb = &message.CallbackURLResponse{}
	case RNowPlaying:
		pb = &message.NowPlayingResponse{}
	case RListen:
		pb = &message.ListenResponse{}
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
