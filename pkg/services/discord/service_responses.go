package discord

import "github.com/bwmarrin/discordgo"

func (s *Service) responseInProgress(
	event *discordgo.InteractionCreate,
	data *CommandData,
) error {
	res := discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Flags: discordgo.MessageFlagsLoading,
		},
	}

	isPrivate := s.isPrivate[data.ID()]
	if isPrivate {
		res.Data.Flags |= discordgo.MessageFlagsEphemeral
	}

	//nolint:wrapcheck
	return s.session.InteractionRespond(event.Interaction, &res)
}

func (s *Service) getNoAccessResponse() *Response {
	return &Response{
		Content: "You have no access to this command",
	}
}

func (s *Service) getNotFoundResponse() *Response {
	return &Response{
		Content: "Command not found",
	}
}
