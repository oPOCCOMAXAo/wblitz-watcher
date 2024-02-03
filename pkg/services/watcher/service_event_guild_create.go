package watcher

import (
	"context"

	"github.com/bwmarrin/discordgo"

	"github.com/opoccomaxao/wblitz-watcher/pkg/services/discord"
	"github.com/opoccomaxao/wblitz-watcher/pkg/services/domain"
)

func (s *Service) eventGuildCreate(
	ctx context.Context,
	event *discord.Event,
) error {
	return s.initNewGuild(ctx, event.Guild)
}

func (s *Service) initNewGuild(
	ctx context.Context,
	guild *discordgo.Guild,
) error {
	err := s.domain.FastFixDiscordGuild(ctx, &domain.FastFixParams{
		ServerID:  guild.ID,
		ChannelID: "",
	})
	if err != nil {
		//nolint:wrapcheck
		return err
	}

	return nil
}
