package discord

import (
	"github.com/bwmarrin/discordgo"
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
					s.getUsernameOption(),
					s.getWotbServerOption(),
				},
			},
		},
		"channelbind": {
			Command: &discordgo.ApplicationCommand{
				Name:        "channelbind",
				Description: "Bind channel for notifications",
				Options: []*discordgo.ApplicationCommandOption{
					{
						Type:        discordgo.ApplicationCommandOptionChannel,
						Name:        "channel",
						Description: "Channel for notifications",
						Required:    true,
					},
					s.getNotificationTypeOption(),
				},
			},
		},
		"clanadd": {
			Command: &discordgo.ApplicationCommand{
				Name:        "clanadd",
				Description: "Add clan for notifications",
				Options: []*discordgo.ApplicationCommandOption{
					s.getClanOption(),
					s.getWotbServerOption(),
				},
			},
		},
		"clanremove": {
			Command: &discordgo.ApplicationCommand{
				Name:        "clanremove",
				Description: "Remove clan from notifications",
				Options: []*discordgo.ApplicationCommandOption{
					s.getClanOption(),
					s.getWotbServerOption(),
				},
			},
		},
		"clanlist": {
			Command: &discordgo.ApplicationCommand{
				Name:        "clanlist",
				Description: "List of clans for notifications",
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
