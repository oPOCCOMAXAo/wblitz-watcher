package watcher

import (
	"github.com/opoccomaxao/wblitz-watcher/pkg/clients/wg"
	"github.com/opoccomaxao/wblitz-watcher/pkg/services/discord"
)

type Service struct {
	discord *discord.Service
	wg      *wg.Client
}

func NewService(
	discord *discord.Service,
	wg *wg.Client,
) *Service {
	return &Service{
		discord: discord,
		wg:      wg,
	}
}

func (s *Service) Serve() error {
	s.discord.RegisterCommandHandler("userstats", s.cmdUserStats)

	return nil
}

func (s *Service) Shutdown() error {
	return nil
}
