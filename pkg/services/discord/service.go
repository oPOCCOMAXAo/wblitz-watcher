package discord

import (
	"github.com/bwmarrin/discordgo"
	"github.com/pkg/errors"

	"github.com/opoccomaxao/wblitz-watcher/pkg/models"
)

type Service struct {
	config   Config
	session  *discordgo.Session
	owner    *discordgo.User
	commands map[string]*InteractionDescription
}

type Config struct {
	ApplicationID string `env:"APPLICATION_ID,required"`
	BotToken      string `env:"BOT_TOKEN,required"`
}

func New(
	config Config,
) (*Service, error) {
	res := Service{
		config: config,
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

	s.commands = s.MakeCommands()

	s.session.AddHandler(s.onReady)
	s.session.AddHandler(s.onInteractionCreate)

	return nil
}
