package domain

import (
	"context"

	"github.com/opoccomaxao/wblitz-watcher/pkg/models"
	"github.com/opoccomaxao/wblitz-watcher/pkg/utils/diff"
)

func (s *Service) CreateEventClanByMembersDiff(
	ctx context.Context,
	values *diff.Diff[*models.WGClanMember],
) (int64, error) {
	events := make([]*models.EventClan, 0, len(values.Created)+len(values.Deleted))

	for _, value := range values.Created {
		events = append(events, &models.EventClan{
			Region:    value.Region,
			ClanID:    value.ClanID,
			AccountID: value.AccountID,
			Type:      models.ETCEnter,
		})
	}

	for _, value := range values.Deleted {
		events = append(events, &models.EventClan{
			Region:    value.Region,
			ClanID:    value.ClanID,
			AccountID: value.AccountID,
			Type:      models.ETCLeave,
		})
	}

	err := s.repo.CreateEventClan(ctx, events...)
	if err != nil {
		//nolint:wrapcheck
		return 0, err
	}

	return int64(len(events)), nil
}
