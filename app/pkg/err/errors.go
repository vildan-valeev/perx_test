package err

import (
	"errors"
	"fmt"
	"strconv"
)

const (
	Unknown  = "unknown error"
	Resource = "need more resource"
	NotFound = "not found"
)

// Error represents an application-specific error. Application errors can be
// unwrapped by the caller to extract out the code & message.
type Error struct {
	// Machine-readable error code.(non-http code)
	Code int
	// Human-readable error message.
	Message string
}

// Error implements the error interface.
func (e *Error) Error() string {
	return fmt.Sprintf("error: code=%d message=%s", e.Code, e.Message)
}

// ErrorCode unwraps an application error and returns its code.
// Non-application errors always return Internal.
func ErrorCode(err error) string {
	var e *Error

	if err == nil {
		return ""
	} else if errors.As(err, &e) {
		return strconv.Itoa(e.Code)
	}

	return Unknown
}

// ErrorMessage unwraps an application error and returns its message.
// Non-application errors always return "internal".
func ErrorMessage(err error) string {
	var e *Error

	if err == nil {
		return ""
	} else if errors.As(err, &e) {
		return e.Message
	}

	return Unknown
}

// Errorf is a helper function to return an Error with a given code and formatted message.
func Errorf(code int, format string, args ...interface{}) *Error {
	return &Error{
		Code:    code,
		Message: fmt.Sprintf(format, args...),
	}
}
