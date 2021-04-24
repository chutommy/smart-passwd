package utils

import (
	"errors"
)

// ErrNilValue is returned whenever variable with nil value is not expected.
var ErrNilValue = errors.New("invalid nil value")
