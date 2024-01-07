package domain

import (
	"context"
	"errors"

	"github.com/opoccomaxao/wblitz-watcher/pkg/models"
)

func (s *Service) CreateUpdateInstance(
	ctx context.Context,
	instance *models.BotInstance,
) error {
	record, err := s.repo.GetInstanceByServer(ctx, instance.ServerID)
	if err != nil && !errors.Is(err, models.ErrNotFound) {
		//nolint:wrapcheck
		return err
	}

	if record == nil {
		err = s.repo.CreateInstance(ctx, instance)
		if err != nil {
			//nolint:wrapcheck
			return err
		}

		return nil
	}

	instance.ID = record.ID

	//nolint:wrapcheck
	return s.repo.UpdateInstance(
		ctx,
		instance,
	)
}
