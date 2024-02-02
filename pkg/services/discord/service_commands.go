package discord

import (
	"context"

	"github.com/bwmarrin/discordgo"

	"github.com/opoccomaxao/wblitz-watcher/pkg/models"
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
	_ context.Context,
	_ *discordgo.InteractionCreate,
	_ *CommandData,
) (*Response, error) {
	return &Response{
		Content: "Pong!",
	}, nil
}

func (s *Service) getServerCommands(
	ctx context.Context,
	guildID string,
) ([]*discordgo.ApplicationCommand, error) {
	//nolint:wrapcheck
	return s.session.ApplicationCommands(
		s.config.ApplicationID,
		guildID,
		s.requestOptions(ctx)...,
	)
}

func (s *Service) parseCommandIDs(
	cmds []*discordgo.ApplicationCommand,
) map[CommandFullName]string {
	res := map[CommandFullName]string{}

	for _, cmd := range cmds {
		res[CommandFullName{Name: cmd.Name}] = cmd.ID

		for _, opt := range cmd.Options {
			if opt.Type != discordgo.ApplicationCommandOptionSubCommand {
				continue
			}

			id := CommandFullName{
				Name:    cmd.Name,
				SubName: opt.Name,
			}

			res[id] = cmd.ID
		}
	}

	return res
}

func (s *Service) processCommand(
	ctx context.Context,
	event *discordgo.InteractionCreate,
	data *CommandData,
) (*Response, error) {
	id := data.ID()

	handler, ok := s.handlers[id]
	if !ok || handler == nil {
		return nil, models.ErrNotFound
	}

	if s.isRestricted[id] {
		err := s.VerifyAccess(event)
		if err != nil {
			return nil, err
		}
	}

	resp, err := handler(ctx, event, data)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
