package configs

import (
	"errors"
)

var (
	// ErrNotFound is key not found.
	ErrNotFound = errors.New("key not found")
	// ErrTypeAssert is type assert error.
	ErrInteralFound = errors.New("internal error")

	ErrParamsInvalid = errors.New("params is invalid")
)
