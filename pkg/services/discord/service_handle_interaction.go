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
	log.Printf("%s\n", data.Name)

	ctx, span := s.tracer.Start(
		context.Background(),
		"cmd:"+data.RequestName(),
		trace.WithSpanKind(trace.SpanKindClient),
	)
	defer span.End()

	s.writeTelemetry(span, event, data)

	err := s.responseInProgress(event, data)
	if err != nil {
		telemetry.RecordErrorFail(ctx, err)
	}

	resp, err := s.processEvent(ctx, event, data)
	if err != nil {
		switch {
		case errors.Is(err, models.ErrNoAccess):
			resp = s.getNoAccessResponse()
		case errors.Is(err, models.ErrNotFound):
			resp = s.getNotFoundResponse()
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

func (s *Service) writeTelemetry(
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

func (s *Service) processEvent(
	ctx context.Context,
	event *discordgo.InteractionCreate,
	data *CommandData,
) (*Response, error) {
	id := data.ID()

	handler, ok := s.handlers[id]
	if !ok || handler == nil {
		return nil, models.ErrNotFound
	}

	if s.isRestricted[id] {
		err := s.VerifyAccess(event)
		if err != nil {
			return nil, err
		}
	}

	resp, err := handler(ctx, event, data)
	if err != nil {
		return nil, err
	}

	return resp, nil
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
