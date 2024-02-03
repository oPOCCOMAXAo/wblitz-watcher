package discord

import (
	"context"
	"net/http"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/pkg/errors"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel/trace"

	"github.com/opoccomaxao/wblitz-watcher/pkg/models"
	"github.com/opoccomaxao/wblitz-watcher/pkg/services/telemetry"
)

type Service struct {
	config            Config
	cmdTracer         trace.Tracer
	eventTracer       trace.Tracer
	client            *http.Client
	session           *discordgo.Session
	handlers          map[CommandFullName]CommandHandler
	eventHandlers     map[EventName]EventHandler
	isRestricted      map[CommandFullName]bool
	isPrivate         map[CommandFullName]bool
	ignoredChannelMap map[string]bool
	useOnlyChannels   bool
	onlyChannelsMap   map[string]bool
	superUserMap      map[string]bool
	existingCommands  map[string]bool
}

type Config struct {
	ApplicationID   string   `env:"APPLICATION_ID,required"`
	BotToken        string   `env:"BOT_TOKEN,required"`
	IgnoreChannels  []string `env:"IGNORE_CHANNELS"`
	UseOnlyChannels []string `env:"USE_ONLY_CHANNELS"`
	SuperUsers      []string `env:"SUPER_USERS"`
}

func New(
	config Config,
	telemetry *telemetry.Service,
) (*Service, error) {
	res := Service{
		config: config,
		cmdTracer: telemetry.PackageTracer("discord",
			trace.WithSpanKind(trace.SpanKindServer),
			models.SpanTypeCommand.Option(),
		),
		eventTracer: telemetry.PackageTracer(
			"discord",
			trace.WithSpanKind(trace.SpanKindConsumer),
			models.SpanTypeEvent.Option(),
		),
		client: &http.Client{
			Timeout:   30 * time.Second,
			Transport: otelhttp.NewTransport(http.DefaultTransport),
		},

		handlers:          map[CommandFullName]CommandHandler{},
		eventHandlers:     map[EventName]EventHandler{},
		isRestricted:      map[CommandFullName]bool{},
		isPrivate:         map[CommandFullName]bool{},
		ignoredChannelMap: map[string]bool{},
		onlyChannelsMap:   map[string]bool{},
		superUserMap:      map[string]bool{},
		existingCommands:  map[string]bool{},
	}

	return &res, res.init()
}

func (s *Service) Serve() error {
	//nolint:wrapcheck
	return s.session.Open()
}

func (s *Service) Shutdown() error {
	//nolint:wrapcheck
	return s.session.Close()
}

func (s *Service) init() error {
	var err error

	s.session, err = discordgo.New("Bot " + s.config.BotToken)
	if err != nil {
		//nolint:wrapcheck
		return err
	}

	application, err := s.session.Application("@me")
	if err != nil {
		//nolint:wrapcheck
		return err
	}

	if application.Owner == nil {
		return errors.WithMessage(models.ErrFailed, "owner is nil")
	}

	for _, channelID := range s.config.IgnoreChannels {
		s.ignoredChannelMap[channelID] = true
	}

	for _, channelID := range s.config.UseOnlyChannels {
		s.onlyChannelsMap[channelID] = true
	}

	s.useOnlyChannels = len(s.onlyChannelsMap) > 0

	for _, userID := range s.config.SuperUsers {
		s.superUserMap[userID] = true
	}

	if application.Owner != nil {
		s.superUserMap[application.Owner.ID] = true
	}

	for _, cmd := range s.getCommands() {
		s.existingCommands[cmd.Name] = true
	}

	s.RegisterCommand(CommandParams{
		Name:      "ping",
		Handler:   s.cmdPing,
		IsPrivate: true,
	})
	s.RegisterCommand(CommandParams{
		Name:      "help",
		Handler:   s.cmdHelp,
		IsPrivate: false,
	})

	s.session.AddHandler(s.onReady)
	s.session.AddHandler(s.onInteractionCreate)
	s.session.AddHandler(s.onGuildCreate)
	s.session.AddHandler(s.onGuildDelete)

	return nil
}

func (s *Service) requestOptions(
	ctx context.Context,
) []discordgo.RequestOption {
	return []discordgo.RequestOption{
		discordgo.WithClient(s.client),
		discordgo.WithContext(ctx),
	}
}
