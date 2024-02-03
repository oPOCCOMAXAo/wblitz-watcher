package discord

import (
	"context"

	"github.com/bwmarrin/discordgo"
	"go.opentelemetry.io/otel/attribute"

	"github.com/opoccomaxao/wblitz-watcher/pkg/services/telemetry"
)

func (s *Service) onGuildCreate(
	_ *discordgo.Session,
	event *discordgo.GuildCreate,
) {
	ctx, span := s.eventTracer.Start(context.Background(), "onGuildCreate")
	defer span.End()

	span.SetAttributes(
		attribute.String("event.guild_id", event.Guild.ID),
	)

	s.registerGuildCommands(ctx, event.Guild.ID)

	err := s.processEvent(ctx, EventGuildCreate, &Event{
		Guild: event.Guild,
	})
	if err != nil {
		telemetry.RecordErrorFail(ctx, err)
	}
}
