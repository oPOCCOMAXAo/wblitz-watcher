package discord

import (
	"github.com/bwmarrin/discordgo"

	"github.com/opoccomaxao/wblitz-watcher/pkg/clients/wg"
)

func (s *Service) MakeCommands() map[string]*InteractionDescription {
	return map[string]*InteractionDescription{
		"ping": {
			Handler: s.Ping,
			Command: &discordgo.ApplicationCommand{
				Name:        "ping",
				Description: "Ping websocket",
			},
		},
		"userstats": {
			Command: &discordgo.ApplicationCommand{
				Name:        "userstats",
				Description: "Get user stats",
				Options: []*discordgo.ApplicationCommandOption{
					{
						Type:        discordgo.ApplicationCommandOptionString,
						Name:        "username",
						Description: "WotBlitz username",
						Required:    true,
					},
					{
						Type:        discordgo.ApplicationCommandOptionString,
						Name:        "server",
						Description: "WotBlitz server",
						Required:    true,
						Choices: []*discordgo.ApplicationCommandOptionChoice{
							{
								Name:  "EU",
								Value: wg.RegionEU.Name(),
							},
							{
								Name:  "NA",
								Value: wg.RegionNA.Name(),
							},
							{
								Name:  "ASIA",
								Value: wg.RegionAsia.Name(),
							},
							{
								Name:  "RU - unsupported now",
								Value: wg.RegionRU.Name(),
							},
						},
					},
				},
			},
		},
	}
}

func (s *Service) RegisterCommandHandler(
	name string,
	handler CommandHandler,
) {
	cmd, ok := s.commands[name]
	if !ok {
		return
	}

	cmd.Handler = handler
}

func (s *Service) Ping(
	_ *discordgo.InteractionCreate,
) (*discordgo.InteractionResponse, error) {
	return &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Pong!",
		},
	}, nil
}
