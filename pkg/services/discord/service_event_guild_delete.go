package discord

import (
	"context"

	"github.com/bwmarrin/discordgo"
	"go.opentelemetry.io/otel/attribute"

	"github.com/opoccomaxao/wblitz-watcher/pkg/services/telemetry"
)

func (s *Service) onGuildDelete(
	_ *discordgo.Session,
	event *discordgo.GuildDelete,
) {
	ctx, span := s.eventTracer.Start(context.Background(), "onGuildDelete")
	defer span.End()

	span.SetAttributes(
		attribute.String("event.guild_id", event.Guild.ID),
	)

	if event.Guild.Unavailable {
		return
	}

	err := s.processEvent(ctx, EventGuildDelete, &Event{
		Guild: event.Guild,
	})
	if err != nil {
		telemetry.RecordErrorFail(ctx, err)
	}
}
