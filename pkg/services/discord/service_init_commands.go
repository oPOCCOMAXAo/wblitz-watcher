package discord

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func (s *Service) onReady(
	_ *discordgo.Session,
	event *discordgo.Ready,
) {
	log.Printf("%+v\n", event)

	bulkCmds := s.getCommands()

	names := map[string]bool{}
	for _, cmd := range bulkCmds {
		names[cmd.Name] = true
	}

	for _, guild := range event.Guilds {
		cmds, err := s.session.ApplicationCommands(s.config.ApplicationID, guild.ID)
		if err != nil {
			log.Printf("%+v\n", err)
		}

		found := map[string]bool{}

		for _, cmd := range cmds {
			_, ok := names[cmd.Name]
			if !ok {
				err = s.session.ApplicationCommandDelete(
					s.config.ApplicationID,
					guild.ID,
					cmd.ID,
				)
				if err != nil {
					log.Printf("%+v\n", err)
				}
			}

			found[cmd.Name] = true
		}

		_, err = s.session.ApplicationCommandBulkOverwrite(s.config.ApplicationID, guild.ID, bulkCmds)
		if err != nil {
			log.Printf("%+v\n", err)
		}
	}
}

//nolint:funlen
func (s *Service) getCommands() []*discordgo.ApplicationCommand {
	return []*discordgo.ApplicationCommand{
		{
			Name:        "ping",
			Description: "Ping websocket",
		},
		{
			Name:        "userstats",
			Description: "Get user stats",
			Options: []*discordgo.ApplicationCommandOption{
				s.getUsernameOption(),
				s.getWotbServerOption(),
			},
		},
		{
			Name:        "user",
			Description: "User commands",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionSubCommand,
					Name:        "stats",
					Description: "Get user stats",
					Options: []*discordgo.ApplicationCommandOption{
						s.getUsernameOption(),
						s.getWotbServerOption(),
					},
				},
			},
		},
		{
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
		{
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
	}
}
