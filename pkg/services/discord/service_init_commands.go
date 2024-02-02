package discord

import (
	"context"

	"github.com/bwmarrin/discordgo"

	"github.com/opoccomaxao/wblitz-watcher/pkg/services/telemetry"
)

func (s *Service) onReady(
	_ *discordgo.Session,
	event *discordgo.Ready,
) {
	ctx, span := s.eventTracer.Start(context.Background(), "onReady")
	defer span.End()

	for _, guild := range event.Guilds {
		s.registerGuildCommands(ctx, guild.ID)
	}
}

func (s *Service) registerGuildCommands(
	ctx context.Context,
	guildID string,
) {
	cmds, err := s.session.ApplicationCommands(
		s.config.ApplicationID,
		guildID,
		s.requestOptions(ctx)...,
	)
	if err != nil {
		telemetry.RecordError(ctx, err)
	}

	found := map[string]bool{}

	for _, cmd := range cmds {
		_, ok := s.existingCommands[cmd.Name]
		if !ok {
			err = s.session.ApplicationCommandDelete(
				s.config.ApplicationID,
				guildID,
				cmd.ID,
				s.requestOptions(ctx)...,
			)
			if err != nil {
				telemetry.RecordError(ctx, err)
			}
		}

		found[cmd.Name] = true
	}

	_, err = s.session.ApplicationCommandBulkOverwrite(
		s.config.ApplicationID,
		guildID,
		s.getCommands(),
		s.requestOptions(ctx)...,
	)
	if err != nil {
		telemetry.RecordError(ctx, err)
	}
}

//nolint:funlen
func (s *Service) getCommands() []*discordgo.ApplicationCommand {
	return []*discordgo.ApplicationCommand{
		{
			Name:        "ping",
			Description: "Ping websocket",
			DescriptionLocalizations: &map[discordgo.Locale]string{
				"uk": "Пінг",
				"ru": "Пинг",
			},
		},
		{
			Name:        "help",
			Description: "Help",
			DescriptionLocalizations: &map[discordgo.Locale]string{
				"uk": "Допомога",
				"ru": "Помощь",
			},
		},
		{
			Name:        "user",
			Description: "User commands",
			DescriptionLocalizations: &map[discordgo.Locale]string{
				"uk": "Команди користувача",
				"ru": "Команды пользователя",
			},
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionSubCommand,
					Name:        "stats",
					Description: "Get user stats",
					DescriptionLocalizations: map[discordgo.Locale]string{
						"uk": "Отримати статистику користувача",
						"ru": "Получить статистику пользователя",
					},
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
			DescriptionLocalizations: &map[discordgo.Locale]string{
				"uk": "Команди каналу",
				"ru": "Команды канала",
			},
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionSubCommand,
					Name:        "bind",
					Description: "Bind channel for notifications",
					DescriptionLocalizations: map[discordgo.Locale]string{
						"uk": "Встановити канал для сповіщень",
						"ru": "Установить канал для уведомлений",
					},
					Options: []*discordgo.ApplicationCommandOption{
						{
							Type:        discordgo.ApplicationCommandOptionChannel,
							Name:        "channel",
							Description: "Channel for notifications",
							DescriptionLocalizations: map[discordgo.Locale]string{
								"uk": "Канал для сповіщень",
								"ru": "Канал для уведомлений",
							},
							Required: true,
						},
						s.getNotificationTypeOption(),
					},
				},
			},
		},
		{
			Name:        "clan",
			Description: "Clan commands",
			DescriptionLocalizations: &map[discordgo.Locale]string{
				"uk": "Команди клану",
				"ru": "Команды клана",
			},
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionSubCommand,
					Name:        "list",
					Description: "List of clans for notifications",
					DescriptionLocalizations: map[discordgo.Locale]string{
						"uk": "Список кланів для сповіщень",
						"ru": "Список кланов для уведомлений",
					},
				},
				{
					Type:        discordgo.ApplicationCommandOptionSubCommand,
					Name:        "add",
					Description: "Add clan for notifications",
					DescriptionLocalizations: map[discordgo.Locale]string{
						"uk": "Додати клан для сповіщень",
						"ru": "Добавить клан для уведомлений",
					},
					Options: []*discordgo.ApplicationCommandOption{
						s.getWotbServerOption(),
						s.getClanOption(),
					},
				},
				{
					Type:        discordgo.ApplicationCommandOptionSubCommand,
					Name:        "remove",
					Description: "Remove clan from notifications",
					DescriptionLocalizations: map[discordgo.Locale]string{
						"uk": "Видалити клан зі сповіщень",
						"ru": "Удалить клан из уведомлений",
					},
					Options: []*discordgo.ApplicationCommandOption{
						s.getWotbServerOption(),
						s.getClanOption(),
					},
				},
			},
		},
	}
}
