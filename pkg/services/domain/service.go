package domain

import (
	"time"

	"github.com/opoccomaxao/wblitz-watcher/pkg/clients/wg"
	"github.com/opoccomaxao/wblitz-watcher/pkg/repo"
	"github.com/opoccomaxao/wblitz-watcher/pkg/services/discord"
)

type Service struct {
	repo    repo.Repository
	wg      *wg.Client
	discord *discord.Service
}

func NewService(
	repo repo.Repository,
	wg *wg.Client,
	discord *discord.Service,
) *Service {
	return &Service{
		repo:    repo,
		wg:      wg,
		discord: discord,
	}
}

func (s *Service) now() int64 {
	return time.Now().Unix()
}
