package domain

import (
	"time"

	"github.com/opoccomaxao/wblitz-watcher/pkg/clients/wg"
	"github.com/opoccomaxao/wblitz-watcher/pkg/repo"
)

type Service struct {
	repo repo.Repository
	wg   *wg.Client
}

func NewService(
	repo repo.Repository,
	wg *wg.Client,
) *Service {
	return &Service{
		repo: repo,
		wg:   wg,
	}
}

func (s *Service) now() int64 {
	return time.Now().Unix()
}
