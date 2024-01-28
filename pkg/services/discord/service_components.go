package discord

import "github.com/bwmarrin/discordgo"

func (s *Service) copyrightFooter() *discordgo.MessageEmbedFooter {
	return &discordgo.MessageEmbedFooter{
		Text:    "Wolverine © 2024",
		IconURL: "https://cdn.discordapp.com/avatars/478551579117748225/6f88a82960e0e236e89d7b7a21e0b6df.webp",
	}
}
