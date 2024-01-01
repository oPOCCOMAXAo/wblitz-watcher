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
	s.discord.RegisterCommandHandler("channelbind", s.cmdChannelBind)
	s.discord.RegisterCommandHandler("clanadd", s.cmdClanAdd)

	return nil
}

func (s *Service) Shutdown() error {
	return nil
}
