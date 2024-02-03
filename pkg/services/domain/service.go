package domain

import (
	"sync"
	"time"

	"github.com/opoccomaxao/wblitz-watcher/pkg/clients/wg"
	"github.com/opoccomaxao/wblitz-watcher/pkg/repo"
	"github.com/opoccomaxao/wblitz-watcher/pkg/services/discord"
)

type Service struct {
	repo    repo.Repository
	wg      *wg.Client
	discord *discord.Service

	fastFixExecutions map[string]struct{}
	fastFixMutex      sync.Mutex
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

		fastFixExecutions: make(map[string]struct{}, 100),
		fastFixMutex:      sync.Mutex{},
	}
}

func (s *Service) now() int64 {
	return time.Now().Unix()
}
