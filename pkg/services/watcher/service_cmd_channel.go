package watcher

import "github.com/bwmarrin/discordgo"

func (s *Service) cmdChannelBind(
	event *discordgo.InteractionCreate,
) (*discordgo.InteractionResponse, error) {
	return &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Debug",
			Embeds: []*discordgo.MessageEmbed{
				{
					Fields: []*discordgo.MessageEmbedField{
						{
							Name:  "Channel",
							Value: event.ChannelID,
						},
						{
							Name:  "Server",
							Value: event.GuildID,
						},
					},
				},
			},
		},
	}, nil
}
