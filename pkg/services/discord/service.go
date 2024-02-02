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
	config           Config
	isProd           bool
	cmdTracer        trace.Tracer
	eventTracer      trace.Tracer
	client           *http.Client
	session          *discordgo.Session
	owner            *discordgo.User
	handlers         map[CommandFullName]CommandHandler
	eventHandlers    map[EventName]EventHandler
	isRestricted     map[CommandFullName]bool
	isPrivate        map[CommandFullName]bool
	isTest           map[string]bool
	existingCommands map[string]bool
}

type Config struct {
	ApplicationID string   `env:"APPLICATION_ID,required"`
	BotToken      string   `env:"BOT_TOKEN,required"`
	TestChannels  []string `env:"TEST_CHANNELS"`
}

func New(
	config Config,
	env models.Environment,
	telemetry *telemetry.Service,
) (*Service, error) {
	res := Service{
		config: config,
		isProd: env.IsProduction(),
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

		handlers:         map[CommandFullName]CommandHandler{},
		eventHandlers:    map[EventName]EventHandler{},
		isRestricted:     map[CommandFullName]bool{},
		isPrivate:        map[CommandFullName]bool{},
		isTest:           map[string]bool{},
		existingCommands: map[string]bool{},
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

	s.owner = application.Owner

	for _, channelID := range s.config.TestChannels {
		s.isTest[channelID] = true
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
