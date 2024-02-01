package watcher

import (
	"context"
)

func (s *Service) serveTasks(ctx context.Context, cancel context.CancelCauseFunc) {
	done := ctx.Done()

	s.execTaskSingle(ctx, cancel, s.TaskSendMessages)
	s.execTaskSingle(ctx, cancel, s.TaskProcessEvents)
	s.execTaskSingle(ctx, cancel, s.TaskWatchClan)

	for {
		select {
		case <-s.tickerWatchClan.C:
			s.execTaskSingle(ctx, cancel, s.TaskWatchClan)
		case <-s.chanProcessEvents:
			s.execTaskSingle(ctx, cancel, s.TaskProcessEvents)
		case <-s.chanSendMessages:
			s.execTaskSingle(ctx, cancel, s.TaskSendMessages)
		case <-done:
			return
		}
	}
}

func (s *Service) execTaskSingle(
	ctx context.Context,
	cancel context.CancelCauseFunc,
	task func(context.Context) error,
) {
	err := task(ctx)
	if err != nil {
		cancel(err)
	}
}

func (s *Service) execInitialTasks(ctx context.Context) error {
	err := s.domain.UpdateAllSubscriptionClanActivation(ctx)
	if err != nil {
		//nolint:wrapcheck
		return err
	}

	return nil
}
