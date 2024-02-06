package discord

import (
	"context"

	"github.com/bwmarrin/discordgo"
)

func (s *Service) SendMessage(
	ctx context.Context,
	channelID string,
	data *Response,
) error {
	_, err := s.session.ChannelMessageSendComplex(
		channelID,
		data.MessageSend(),
		discordgo.WithContext(ctx),
		discordgo.WithClient(s.client),
	)
	if err != nil {
		return MapError(err)
	}

	return nil
}
