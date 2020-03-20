package types

// InputType …
type InputType byte

// InputTypeEnum …
const (
	IPing        InputType = 'p'
	IAction      InputType = 'a'
	INewUser     InputType = 'n'
	IUserExists  InputType = 'e'
	ICallbackURL InputType = 'c'
)

// Inv Response type
func (i InputType) Inv() (r ResponseType) {
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
	}
	return
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
)

// Inv Input type
func (r ResponseType) Inv() (i InputType) {
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
	}
	return
}
