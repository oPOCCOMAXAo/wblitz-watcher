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

	total, err := s.taskProcessEvents(ctx)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
	}

	if total > 0 {
		s.notifyTaskSendMessages()
	}

	return err
}

func (s *Service) taskProcessEvents(
	ctx context.Context,
) (int64, error) {
	total, err := s.domain.CreateDiscordMessagesForEventClan(ctx)
	if err != nil {
		//nolint:wrapcheck
		return total, err
	}

	return total, nil
}
