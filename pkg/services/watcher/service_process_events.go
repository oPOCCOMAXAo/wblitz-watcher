package watcher

import (
	"context"

	"go.opentelemetry.io/otel/codes"
)

func (s *Service) notifyTaskProcessEvents() {
	select {
	case s.chanProcessEvents <- struct{}{}:
	default:
	}
}

func (s *Service) TaskProcessEvents(
	ctx context.Context,
) error {
	ctx, span := s.taskTracer.Start(ctx, "processEvents")
	defer span.End()

	err := s.taskProcessEvents(ctx)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
	}

	s.notifyTaskSendMessages()

	return err
}

func (s *Service) taskProcessEvents(
	ctx context.Context,
) error {
	err := s.domain.CreateDiscordMessagesForEventClan(ctx)
	if err != nil {
		//nolint:wrapcheck
		return err
	}

	return nil
}
