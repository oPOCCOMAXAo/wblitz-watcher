package watcher

import (
	"context"

	"github.com/opoccomaxao/wblitz-watcher/pkg/services/discord"
)

func (s *Service) eventGuildDelete(
	ctx context.Context,
	event *discord.Event,
) error {
	err := s.domain.DeleteDiscordGuildData(ctx, event.Guild.ID)
	if err != nil {
		//nolint:wrapcheck
		return err
	}

	return nil
}
