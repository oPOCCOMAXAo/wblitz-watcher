package wg

import "errors"

var (
	ErrAPI           = errors.New("api error")
	ErrLimitExceeded = errors.New("limit exceeded")
)
