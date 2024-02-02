package discord

import (
	"context"

	"github.com/bwmarrin/discordgo"
	"go.opentelemetry.io/otel/attribute"
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
}
