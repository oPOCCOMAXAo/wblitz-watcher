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

	err := s.processEvent(ctx, EventReady, &Event{
		Guilds: event.Guilds,
	})
	if err != nil {
		telemetry.RecordErrorFail(ctx, err)
	}
}
