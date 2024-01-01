package discord

import (
	"github.com/bwmarrin/discordgo"
)

type Service struct {
	config   Config
	session  *discordgo.Session
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

	s.commands = s.MakeCommands()

	s.session.AddHandler(s.onReady)
	s.session.AddHandler(s.onInteractionCreate)

	return nil
}
