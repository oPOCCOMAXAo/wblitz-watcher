package watcher

import (
	"github.com/bwmarrin/discordgo"

	"github.com/opoccomaxao/wblitz-watcher/pkg/app"
	"github.com/opoccomaxao/wblitz-watcher/pkg/models"
	du "github.com/opoccomaxao/wblitz-watcher/pkg/utils/discordutils"
)

func (s *Service) cmdChannel(
	event *discordgo.InteractionCreate,
) (*discordgo.InteractionResponse, error) {
	data := event.ApplicationCommandData()

	switch data.Options[0].Name {
	case "bind":
		return s.cmdChannelBind(event)
	}

	return nil, app.ErrNotFound
}

func (s *Service) cmdChannelBind(
	event *discordgo.InteractionCreate,
) (*discordgo.InteractionResponse, error) {
	err := s.discord.VerifyAccess(event)
	if err != nil {
		//nolint:wrapcheck
		return nil, err
	}

	var instance models.BotInstance

	err = du.ParseOptions(event.ApplicationCommandData().Options[0].Options, du.DecodersMap{
		"channel": du.DecoderChannelID(&instance.ChannelID),
		"type":    du.DecoderString(&instance.Type),
	})
	if err != nil {
		//nolint:wrapcheck
		return nil, err
	}

	instance.ServerID = event.GuildID

	return &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Debug",
			Embeds: []*discordgo.MessageEmbed{
				{
					Fields: []*discordgo.MessageEmbedField{
						{
							Name:  "Channel",
							Value: instance.ChannelID,
						},
						{
							Name:  "Server",
							Value: instance.ServerID,
						},
						{
							Name:  "Type",
							Value: string(instance.Type),
						},
					},
				},
			},
			Flags: discordgo.MessageFlagsEphemeral,
		},
	}, nil
}
