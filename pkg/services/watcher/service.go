package watcher

import (
	"github.com/samber/do"

	"github.com/opoccomaxao/wblitz-watcher/pkg/clients/wg"
	"github.com/opoccomaxao/wblitz-watcher/pkg/services/discord"
	"github.com/opoccomaxao/wblitz-watcher/pkg/services/domain"
)

var _ do.Shutdownable = (*Service)(nil)

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
		Name:    "user",
		SubName: "stats",
		Handler: s.cmdUserStats,
	})

	s.discord.RegisterCommand(discord.CommandParams{
		Name:         "channel",
		SubName:      "bind",
		Handler:      s.cmdChannelBind,
		IsRestricted: true,
		IsPrivate:    true,
	})

	s.discord.RegisterCommand(discord.CommandParams{
		Name:         "clan",
		SubName:      "add",
		Handler:      s.cmdClanAdd,
		IsRestricted: true,
		IsPrivate:    true,
	})

	s.discord.RegisterCommand(discord.CommandParams{
		Name:         "clan",
		SubName:      "remove",
		Handler:      s.cmdClanRemove,
		IsRestricted: true,
		IsPrivate:    true,
	})

	s.discord.RegisterCommand(discord.CommandParams{
		Name:    "clan",
		SubName: "list",
		Handler: s.cmdClanList,
	})

	return nil
}

func (s *Service) Shutdown() error {
	return nil
}
