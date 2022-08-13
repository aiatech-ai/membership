package errs

import (
	"fmt"
	"runtime"
)

type PrivateError struct {
	Message string
	Err     error
	Caller  string
}

type PublicError struct {
	Message string
	Caller  string
}

func (s PrivateError) Error() string {
	return fmt.Sprintf("%s, detail is: %s\n", s.Message, s.Err.Error())
}
func (s PublicError) Error() string {
	return s.Message
}

var (
	ErrInvalidCredential = "invalid credential"
	ErrUserNotFound      = "user not found"
	ErrUserAlreadyExists = "user already exists"
	UnablePrepareQuery   = "unable prepare query"
)

// NewPrivateError: contain origin error from system
// message: final message (include value interface{})
// err: origin error
func NewPrivateError(message string, err error) PrivateError {
	pc, file, line, _ := runtime.Caller(1)
	details := runtime.FuncForPC(pc)

	return PrivateError{
		Message: message,
		Err:     err,
		Caller:  fmt.Sprintf("%s#%s#%d", file, details.Name(), line),
	}
}

// NewPrivateError
func NewPublicError(message string) PublicError {
	pc, file, line, _ := runtime.Caller(1)
	details := runtime.FuncForPC(pc)

	return PublicError{
		Message: message,
		Caller:  fmt.Sprintf("%s#%s#%d", file, details.Name(), line),
	}
}
