package types

// InputType …
type InputType byte

// InputTypeEnum …
const (
	IPing        InputType = 'p'
	IAction      InputType = 'a'
	INewUser     InputType = 'n'
	IUserExist   InputType = 'e'
	ICallbackURL InputType = 'c'
)

// Byte Input byte
func (i *InputType) Byte() (m byte) {
	switch i {
	case IPing:
		m = 'p'
	case IAction:
		m = 'a'
	case INewUser:
		m = 'n'
	case IUserExist:
		m = 'e'
	case ICallbackURL:
		m = 'c'
	}
	return
}

// ResByte Response byte
func (i *InputType) ResByte() (m byte) {
	switch i {
	case IPing:
		m = 'P'
	case IAction:
		m = 'A'
	case INewUser:
		m = 'N'
	case IUserExist:
		m = 'E'
	case ICallbackURL:
		m = 'C'
	}
	return
}

// Res Response type
func (i *InputType) Res() (r ResponseType) {
	switch i {
	case IPing:
		r = RPong
	case IAction:
		r = RAction
	case INewUser:
		r = RNewUser
	case IUserExist:
		r = RUserExists
	case ICallbackURL
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

// Byte Response byte
func (i *ResponseType) Byte() (m byte) {
	switch i {
	case RPong:
		m = 'P'
	case RAction:
		m = 'A'
	case RNewUser:
		m = 'N'
	case RUserExists:
		m = 'E'
	case RCallbackURL:
		m = 'C'
	}
	return
}

// Inp Input type
func (i *ResponseType) Inp() (i InputType) {
	switch i {
	case RPong:
		i = IPing
	case RAction:
		i = IAction
	case RNewUser:
		i = INewUser
	case RUserExists:
		i = IUserExist
	case RCallbackURL:
		i = ICallbackURL
	}
	return
}
