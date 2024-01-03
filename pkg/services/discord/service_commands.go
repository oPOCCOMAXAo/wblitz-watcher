package discord

import (
	"github.com/bwmarrin/discordgo"
)

func (s *Service) MakeCommands() map[string]*InteractionDescription {
	res := map[string]*InteractionDescription{}
	for _, cmd := range s.getCommands() {
		res[cmd.Command.Name] = cmd
	}

	return res
}

//nolint:funlen
func (s *Service) getCommands() []*InteractionDescription {
	return []*InteractionDescription{
		{
			Handler: s.cmdPing,
			Command: &discordgo.ApplicationCommand{
				Name:        "ping",
				Description: "Ping websocket",
			},
		},
		{
			Command: &discordgo.ApplicationCommand{
				Name:        "userstats",
				Description: "Get user stats",
				Options: []*discordgo.ApplicationCommandOption{
					s.getUsernameOption(),
					s.getWotbServerOption(),
				},
			},
		},
		{
			Command: &discordgo.ApplicationCommand{
				Name:        "channel",
				Description: "Channel commands",
				Options: []*discordgo.ApplicationCommandOption{
					{
						Type:        discordgo.ApplicationCommandOptionSubCommand,
						Name:        "bind",
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
			},
		},
		{
			Command: &discordgo.ApplicationCommand{
				Name:        "clan",
				Description: "Clan commands",
				Options: []*discordgo.ApplicationCommandOption{
					{
						Type:        discordgo.ApplicationCommandOptionSubCommand,
						Name:        "list",
						Description: "List of clans for notifications",
					},
					{
						Type:        discordgo.ApplicationCommandOptionSubCommand,
						Name:        "add",
						Description: "Add clan for notifications",
						Options: []*discordgo.ApplicationCommandOption{
							s.getWotbServerOption(),
							s.getClanOption(),
						},
					},
					{
						Type:        discordgo.ApplicationCommandOptionSubCommand,
						Name:        "remove",
						Description: "Remove clan from notifications",
						Options: []*discordgo.ApplicationCommandOption{
							s.getWotbServerOption(),
							s.getClanOption(),
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

func (s *Service) cmdPing(
	_ *discordgo.InteractionCreate,
) (*discordgo.InteractionResponse, error) {
	return &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Pong!",
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	}, nil
}

func (s *Service) getNoAccessResponse() *discordgo.InteractionResponse {
	return &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "You have no access to this command",
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	}
}

func (s *Service) getNotFoundResponse() *discordgo.InteractionResponse {
	return &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Command not found",
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	}
}
