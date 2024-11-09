package watcher

import (
	"context"

	"github.com/opoccomaxao/wblitz-watcher/pkg/services/telemetry"
)

func (s *Service) serveTasks(
	ctx context.Context,
	_ context.CancelCauseFunc,
) {
	done := ctx.Done()

	s.execTaskSingle(ctx, s.TaskSendMessages)
	s.execTaskSingle(ctx, s.TaskProcessEvents)
	s.execTaskSingle(ctx, s.TaskWatchClan)

	for {
		select {
		case <-s.tickerWatchClan.C:
			s.execTaskSingle(ctx, s.TaskWatchClan)
		case <-s.chanProcessEvents:
			s.execTaskSingle(ctx, s.TaskProcessEvents)
		case <-s.chanSendMessages:
			s.execTaskSingle(ctx, s.TaskSendMessages)
		case <-done:
			return
		}
	}
}

func (s *Service) execTaskSingle(
	ctx context.Context,
	task func(context.Context) error,
) {
	err := task(ctx)
	if err != nil {
		telemetry.RecordErrorBackground(err)
	}

	return
}

func (s *Service) execInitialTasks(ctx context.Context) error {
	err := s.domain.UpdateAllSubscriptionClanActivation(ctx)
	if err != nil {
		//nolint:wrapcheck
		return err
	}

	return nil
}
