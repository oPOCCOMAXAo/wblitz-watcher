package domain

import (
	"context"

	"github.com/opoccomaxao/wblitz-watcher/pkg/models"
)

func (s *Service) EnsureWGClan(
	ctx context.Context,
	value *models.WGClan,
) error {
	err := s.repo.CreateUpdateWGClan(ctx, value)
	if err != nil {
		//nolint:wrapcheck
		return err
	}

	return nil
}
