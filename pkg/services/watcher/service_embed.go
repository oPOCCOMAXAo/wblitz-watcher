package watcher

import "github.com/bwmarrin/discordgo"

func (s *Service) embedError(
	err error,
) *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Title: "Error",
		Type:  discordgo.EmbedTypeRich,
		Color: int(ColorError),
	}
}
