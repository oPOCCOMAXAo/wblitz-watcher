package discord

import (
	"github.com/bwmarrin/discordgo"

	"github.com/opoccomaxao/wblitz-watcher/pkg/models"
)

func (s *Service) getWotbServerOption() *discordgo.ApplicationCommandOption {
	return &discordgo.ApplicationCommandOption{
		Type:        discordgo.ApplicationCommandOptionString,
		Name:        "server",
		Description: "WotBlitz server",
		Required:    true,
		Choices: []*discordgo.ApplicationCommandOptionChoice{
			{
				Name:  "EU",
				Value: models.RegionEU,
			},
			{
				Name:  "NA",
				Value: models.RegionNA,
			},
			{
				Name:  "ASIA",
				Value: models.RegionAsia,
			},
			{
				Name:  "RU - unsupported now",
				Value: models.RegionRU,
			},
		},
	}
}

func (s *Service) getUsernameOption() *discordgo.ApplicationCommandOption {
	return &discordgo.ApplicationCommandOption{
		Type:        discordgo.ApplicationCommandOptionString,
		Name:        "username",
		Description: "WotBlitz username",
		Required:    true,
	}
}

func (s *Service) getClanOption() *discordgo.ApplicationCommandOption {
	return &discordgo.ApplicationCommandOption{
		Type:        discordgo.ApplicationCommandOptionString,
		Name:        "clan",
		Description: "Clan tag",
		Required:    true,
	}
}

func (s *Service) getNotificationTypeOption() *discordgo.ApplicationCommandOption {
	return &discordgo.ApplicationCommandOption{
		Type:        discordgo.ApplicationCommandOptionString,
		Name:        "type",
		Description: "Notification type",
		Required:    true,
		Choices: []*discordgo.ApplicationCommandOptionChoice{
			{
				Name:  "Clan notifications",
				Value: "clan",
			},
		},
	}
}
