package watcher

import (
	"context"

	"github.com/opoccomaxao/wblitz-watcher/pkg/services/discord"
)

func (s *Service) eventReady(
	ctx context.Context,
	event *discord.Event,
) error {
	guilds := make([]string, 0, len(event.Guilds))

	for _, guild := range event.Guilds {
		guilds = append(guilds, guild.ID)
	}

	err := s.domain.DeleteDiscordGuildsNotInList(ctx, guilds)
	if err != nil {
		//nolint:wrapcheck
		return err
	}

	for _, guild := range event.Guilds {
		err = s.initNewGuild(ctx, guild)
		if err != nil {
			return err
		}
	}

	return nil
}
