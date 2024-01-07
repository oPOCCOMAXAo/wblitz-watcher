package discord

import "github.com/bwmarrin/discordgo"

func (s *Service) responseInProgress(
	interaction *discordgo.Interaction,
) error {
	return s.session.InteractionRespond(
		interaction,
		&discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Flags: discordgo.MessageFlagsEphemeral,
			},
		},
	)
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
