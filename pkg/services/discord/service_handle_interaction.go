package discord

import (
	"errors"
	"log"

	"github.com/bwmarrin/discordgo"

	"github.com/opoccomaxao/wblitz-watcher/pkg/models"
)

func (s *Service) onInteractionCreate(
	session *discordgo.Session,
	event *discordgo.InteractionCreate,
) {
	err := s.responseInProgress(event.Interaction)
	if err != nil {
		log.Printf("%+v\n", err)
	}

	resp, err := s.processEvent(event)
	if err != nil {
		switch {
		case errors.Is(err, models.ErrNoAccess):
			resp = s.getNoAccessResponse()
		case errors.Is(err, models.ErrNotFound):
			resp = s.getNotFoundResponse()
		default:
			log.Printf("%+v\n", err)
		}
	}

	if resp == nil {
		return
	}

	_, err = session.InteractionResponseEdit(
		event.Interaction,
		resp.WebHookEdit(),
	)
	if err != nil {
		log.Printf("%+v\n", err)
	}
}

func (s *Service) processEvent(
	event *discordgo.InteractionCreate,
) (*Response, error) {
	data := s.parseInteractionData(event)

	log.Printf("%s\n", data.Name)

	id := data.ID()

	handler, ok := s.handlers[id]
	if !ok || handler == nil {
		return nil, models.ErrNotFound
	}

	_, ok = s.accessRequired[id]
	if ok {
		err := s.VerifyAccess(event)
		if err != nil {
			return nil, err
		}
	}

	resp, err := handler(event, data)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *Service) parseInteractionData(
	event *discordgo.InteractionCreate,
) *CommandData {
	data := event.ApplicationCommandData()

	res := CommandData{
		Name:    []string{data.Name},
		Options: map[string]any{},
	}

	for _, opt := range data.Options {
		s.parseOptionInto(opt, &res)
	}

	return &res
}

func (s *Service) parseOptionInto(
	opt *discordgo.ApplicationCommandInteractionDataOption,
	res *CommandData,
) {
	switch opt.Type {
	case discordgo.ApplicationCommandOptionSubCommand:
		res.Name = append(res.Name, opt.Name)
	default:
		res.Options[opt.Name] = opt.Value
	}

	for _, opt := range opt.Options {
		s.parseOptionInto(opt, res)
	}
}
