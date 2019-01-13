package gfx

import "fmt"

// ErrDone can for example be returned when you are done rendering.
var ErrDone = Error("done")

// Error type
type Error string

// Error implements the error interface
func (e Error) Error() string {
	return string(e)
}

// Errorf constructs a formatted error
func Errorf(format string, a ...interface{}) error {
	return Error(fmt.Sprintf(format, a...))
}