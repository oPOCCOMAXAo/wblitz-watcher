package wg

import (
	"errors"

	"github.com/opoccomaxao/wblitz-watcher/pkg/models"
)

var (
	ErrAPI           = errors.New("api error")
	ErrLimitExceeded = errors.New("limit exceeded")
	ErrRetryLater    = models.ErrRetryLater
)
