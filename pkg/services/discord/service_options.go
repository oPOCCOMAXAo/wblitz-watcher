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
		DescriptionLocalizations: map[discordgo.Locale]string{
			"uk": "Сервер WotBlitz",
			"ru": "Сервер WotBlitz",
		},
		Required: true,
		Choices: []*discordgo.ApplicationCommandOptionChoice{
			{
				Name:  "EU - Europe",
				Value: models.RegionEU,
				NameLocalizations: map[discordgo.Locale]string{
					"uk": "EU - Європа",
					"ru": "EU - Европа",
				},
			},
			{
				Name:  "NA - North America",
				Value: models.RegionNA,
				NameLocalizations: map[discordgo.Locale]string{
					"uk": "NA - Північна Америка",
					"ru": "NA - Северная Америка",
				},
			},
			{
				Name:  "ASIA - Asia",
				Value: models.RegionAsia,
				NameLocalizations: map[discordgo.Locale]string{
					"uk": "ASIA - Азія",
					"ru": "ASIA - Азия",
				},
			},
		},
	}
}

func (s *Service) getUsernameOption() *discordgo.ApplicationCommandOption {
	return &discordgo.ApplicationCommandOption{
		Type:        discordgo.ApplicationCommandOptionString,
		Name:        "username",
		Description: "WotBlitz username",
		DescriptionLocalizations: map[discordgo.Locale]string{
			"uk": "Ім'я користувача WotBlitz",
			"ru": "Имя пользователя WotBlitz",
		},
		Required: true,
	}
}

func (s *Service) getClanOption() *discordgo.ApplicationCommandOption {
	return &discordgo.ApplicationCommandOption{
		Type:        discordgo.ApplicationCommandOptionString,
		Name:        "clan",
		Description: "Clan tag",
		DescriptionLocalizations: map[discordgo.Locale]string{
			"uk": "Тег клану",
			"ru": "Тег клана",
		},
		Required: true,
	}
}

func (s *Service) getNotificationTypeOption() *discordgo.ApplicationCommandOption {
	return &discordgo.ApplicationCommandOption{
		Type:        discordgo.ApplicationCommandOptionString,
		Name:        "type",
		Description: "Notification type",
		DescriptionLocalizations: map[discordgo.Locale]string{
			"uk": "Тип сповіщень",
			"ru": "Тип уведомлений",
		},
		Required: true,
		Choices: []*discordgo.ApplicationCommandOptionChoice{
			{
				Name:  "Clan notifications",
				Value: models.STClan,
				NameLocalizations: map[discordgo.Locale]string{
					"uk": "Сповіщення клану",
					"ru": "Уведомления клана",
				},
			},
			{
				Name:  "Info notifications",
				Value: models.STInfo,
				NameLocalizations: map[discordgo.Locale]string{
					"uk": "Інформаційні сповіщення",
					"ru": "Информационные уведомления",
				},
			},
		},
	}
}
