package migrations

import "errors"

var (
	ErrTooShortID        = errors.New("too short id")
	ErrInvalidTimeFormat = errors.New("invalid time format")
	ErrInvalidDate       = errors.New("invalid date")
)
