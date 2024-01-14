package discord

import (
	"github.com/bwmarrin/discordgo"
)

type CommandParams struct {
	Name         string
	SubName      string
	Handler      CommandHandler
	IsRestricted bool
	IsPrivate    bool
}

func (s *Service) RegisterCommand(params CommandParams) {
	id := CommandFullName{
		Name:    params.Name,
		SubName: params.SubName,
	}

	if params.Handler != nil {
		s.handlers[id] = params.Handler
	}

	if params.IsRestricted {
		s.isRestricted[id] = true
	}

	if params.IsPrivate {
		s.isPrivate[id] = true
	}
}

func (s *Service) cmdPing(
	_ *discordgo.InteractionCreate,
	_ *CommandData,
) (*Response, error) {
	return &Response{
		Content: "Pong!",
	}, nil
}
