package models

import "errors"

var (
	ErrAlreadyExists = errors.New("already exists")
	ErrFailed        = errors.New("failed")
	ErrFlowBroken    = errors.New("flow broken")
	ErrNoAccess      = errors.New("no access")
	ErrNotFound      = errors.New("not found")
	ErrPanic         = errors.New("panic")
	ErrRetryLater    = errors.New("retry later")
)
