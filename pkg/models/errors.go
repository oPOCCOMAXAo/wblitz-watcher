package models

import "errors"

var (
	ErrAlreadyExists = errors.New("already exists")
	ErrFailed        = errors.New("failed")
	ErrNoAccess      = errors.New("no access")
	ErrNotFound      = errors.New("not found")
	ErrPanic         = errors.New("panic")
)
