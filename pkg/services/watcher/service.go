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
	s.discord.RegisterCommand(discord.CommandParams{
		Name:    "userstats",
		Handler: s.cmdUserStats,
	})

	s.discord.RegisterCommand(discord.CommandParams{
		Name:      "channel",
		SubName:   "bind",
		Handler:   s.cmdChannelBind,
		IsPrivate: true,
	})

	s.discord.RegisterCommand(discord.CommandParams{
		Name:      "clan",
		SubName:   "add",
		Handler:   s.cmdClanAdd,
		IsPrivate: true,
	})

	s.discord.RegisterCommand(discord.CommandParams{
		Name:      "clan",
		SubName:   "remove",
		Handler:   s.cmdClanRemove,
		IsPrivate: true,
	})

	return nil
}

func (s *Service) Shutdown() error {
	return nil
}
