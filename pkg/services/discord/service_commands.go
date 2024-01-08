package discord

import (
	"github.com/bwmarrin/discordgo"
)

type CommandParams struct {
	Name      string
	SubName   string
	Handler   CommandHandler
	IsPrivate bool
}

func (s *Service) RegisterCommand(params CommandParams) {
	id := CommandFullName{
		Name:    params.Name,
		SubName: params.SubName,
	}

	if params.Handler != nil {
		s.handlers[id] = params.Handler
	}

	if params.IsPrivate {
		s.accessRequired[id] = true
	}
}

func (s *Service) RegisterCommandHandler(
	name string,
	handler CommandHandler,
) {
	s.RegisterCommand(CommandParams{
		Name:    name,
		Handler: handler,
	})
}

func (s *Service) RegisterSubCommandHandler(
	name string,
	subName string,
	handler CommandHandler,
) {
	s.RegisterCommand(CommandParams{
		Name:    name,
		SubName: subName,
		Handler: handler,
	})
}

func (s *Service) cmdPing(
	_ *discordgo.InteractionCreate,
	_ *CommandData,
) (*Response, error) {
	return &Response{
		Content: "Pong!",
	}, nil
}
