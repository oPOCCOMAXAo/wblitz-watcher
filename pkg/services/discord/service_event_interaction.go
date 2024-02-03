package discord

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"

	"github.com/opoccomaxao/wblitz-watcher/pkg/models"
	"github.com/opoccomaxao/wblitz-watcher/pkg/services/telemetry"
)

func (s *Service) onInteractionCreate(
	session *discordgo.Session,
	event *discordgo.InteractionCreate,
) {
	data := s.parseInteractionData(event)

	if s.isChannelIgnored(event.ChannelID) {
		log.Printf("ignored: %s\n", data.Name)

		return
	}

	log.Printf("%s\n", data.Name)

	ctx, span := s.cmdTracer.Start(context.Background(), data.RequestName())
	defer span.End()

	s.writeInteractionTelemetry(span, event, data)

	err := s.responseInProgress(event, data)
	if err != nil {
		telemetry.RecordErrorFail(ctx, err)
	}

	resp, err := s.processCommand(ctx, event, data)
	if err != nil {
		switch {
		case errors.Is(err, models.ErrNoAccess):
			resp = s.getNoAccessResponse()

			telemetry.RecordError(ctx, err)
		case errors.Is(err, models.ErrNotFound):
			resp = s.getNotFoundResponse()

			telemetry.RecordError(ctx, err)
		default:
			telemetry.RecordErrorFail(ctx, err)
		}
	}

	if resp == nil {
		return
	}

	_, err = session.InteractionResponseEdit(
		event.Interaction,
		resp.WebHookEdit(),
	)
	if err != nil {
		telemetry.RecordErrorFail(ctx, err)
	}
}

func (s *Service) writeInteractionTelemetry(
	span trace.Span,
	event *discordgo.InteractionCreate,
	data *CommandData,
) {
	attrs := []attribute.KeyValue{
		attribute.String("interaction.id", event.ID),
		attribute.String("interaction.type", event.Type.String()),
	}

	for key, value := range data.Options {
		attrs = append(attrs, attribute.String(
			"interaction.option."+key,
			fmt.Sprint(value)),
		)
	}

	span.SetAttributes(attrs...)
}

func (s *Service) parseInteractionData(
	event *discordgo.InteractionCreate,
) *CommandData {
	data := event.ApplicationCommandData()

	res := CommandData{
		Name:    []string{data.Name},
		Options: map[string]any{},
	}

	for _, opt := range data.Options {
		s.parseOptionInto(opt, &res)
	}

	return &res
}

func (s *Service) parseOptionInto(
	opt *discordgo.ApplicationCommandInteractionDataOption,
	res *CommandData,
) {
	switch opt.Type {
	case discordgo.ApplicationCommandOptionSubCommand:
		res.Name = append(res.Name, opt.Name)
	default:
		res.Options[opt.Name] = opt.Value
	}

	for _, opt := range opt.Options {
		s.parseOptionInto(opt, res)
	}
}
