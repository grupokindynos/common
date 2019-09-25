package errors

import "errors"

var (
	ErrorNoAuth            = errors.New("you are not authorized")
	ErrorUnmarshal         = errors.New("unable to unmarshal object")
)