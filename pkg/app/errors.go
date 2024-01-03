package app

import "errors"

// errors
var (
	ErrFailed   = errors.New("failed")
	ErrNoAccess = errors.New("no access")
	ErrNotFound = errors.New("not found")
	ErrPanic    = errors.New("panic")
)
