package discord

import "github.com/bwmarrin/discordgo"

func (s *Service) isEventIgnored(
	_ *discordgo.InteractionCreate,
	data *CommandData,
) bool {
	eventIsProd := !data.IsTest

	return s.isProd != eventIsProd
}
