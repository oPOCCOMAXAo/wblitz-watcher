package watcher

import (
	"context"
	"time"

	"github.com/samber/do"
	"go.opentelemetry.io/otel/trace"

	"github.com/opoccomaxao/wblitz-watcher/pkg/clients/wg"
	"github.com/opoccomaxao/wblitz-watcher/pkg/services/discord"
	"github.com/opoccomaxao/wblitz-watcher/pkg/services/domain"
	"github.com/opoccomaxao/wblitz-watcher/pkg/services/telemetry"
)

var _ do.Shutdownable = (*Service)(nil)

type Service struct {
	taskTracer trace.Tracer
	discord    *discord.Service
	wg         *wg.Client
	domain     *domain.Service

	tickerWatchClan   *time.Ticker
	chanSendMessages  chan struct{}
	chanProcessEvents chan struct{}
}

func NewService(
	telemetry *telemetry.Service,
	discord *discord.Service,
	wg *wg.Client,
	domain *domain.Service,
) *Service {
	return &Service{
		taskTracer: telemetry.TaskTracer(),
		discord:    discord,
		wg:         wg,
		domain:     domain,

		tickerWatchClan:   time.NewTicker(1 * time.Minute),
		chanSendMessages:  make(chan struct{}, 10),
		chanProcessEvents: make(chan struct{}, 10),
	}
}

func (s *Service) now() int64 {
	return time.Now().Unix()
}

func (s *Service) Serve(
	ctx context.Context,
	cancel context.CancelCauseFunc,
) error {
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

	s.discord.RegisterEvent(discord.EventParams{
		Name:    discord.EventGuildDelete,
		Handler: s.eventGuildDelete,
	})

	go s.serveTasks(ctx, cancel)

	err := s.execInitialTasks(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) Shutdown() error {
	return nil
}
