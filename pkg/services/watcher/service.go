package watcher

import (
	"github.com/opoccomaxao/wblitz-watcher/pkg/clients/wg"
	"github.com/opoccomaxao/wblitz-watcher/pkg/services/discord"
	"github.com/opoccomaxao/wblitz-watcher/pkg/services/domain"
)

type Service struct {
	discord *discord.Service
	wg      *wg.Client
	domain  *domain.Service
}

func NewService(
	discord *discord.Service,
	wg *wg.Client,
	domain *domain.Service,
) *Service {
	return &Service{
		discord: discord,
		wg:      wg,
		domain:  domain,
	}
}

func (s *Service) Serve() error {
	s.discord.RegisterCommandHandler("userstats", s.cmdUserStats)
	s.discord.RegisterCommandHandler("channel", s.cmdChannel)
	s.discord.RegisterCommandHandler("clan", s.cmdClan)

	return nil
}

func (s *Service) Shutdown() error {
	return nil
}
