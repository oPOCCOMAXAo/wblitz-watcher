package watcher

import (
	"context"

	"github.com/bwmarrin/discordgo"

	"github.com/opoccomaxao/wblitz-watcher/pkg/models"
	"github.com/opoccomaxao/wblitz-watcher/pkg/services/discord"
	"github.com/opoccomaxao/wblitz-watcher/pkg/services/domain"
	"github.com/opoccomaxao/wblitz-watcher/pkg/services/telemetry"
)

func (s *Service) cmdUserStats(
	ctx context.Context,
	event *discordgo.InteractionCreate,
	data *discord.CommandData,
) (*discord.Response, error) {
	request := domain.UserStatsRequest{
		Username: data.String("username"),
		Region:   models.Region(data.String("server")),
	}

	defer s.domain.FastFixDiscordGuildOnce(ctx, &domain.FastFixParams{
		ServerID:  event.GuildID,
		ChannelID: event.ChannelID,
	})

	user, err := s.domain.UserStats(ctx, &request)
	if err != nil {
		telemetry.RecordError(ctx, err)
	}

	if user == nil {
		return &discord.Response{
			Content: "User not found",
		}, nil
	}

	if err != nil {
		//nolint:wrapcheck
		return nil, err
	}

	return &discord.Response{
		Content: "User stats",
		Embeds: []*discordgo.MessageEmbed{
			s.embedAccountInfo(user),
		},
	}, nil
}
