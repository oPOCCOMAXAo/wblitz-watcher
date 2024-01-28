package domain

import (
	"context"
	"errors"

	"github.com/opoccomaxao/wblitz-watcher/pkg/clients/wg"
	"github.com/opoccomaxao/wblitz-watcher/pkg/models"
)

type UserStatsRequest struct {
	ServerID string
	Username string
	Region   models.Region
}

func (s *Service) UserStats(
	ctx context.Context,
	request *UserStatsRequest,
) (*wg.AccountInfo, error) {
	listEntry, err := s.wg.FindAccountByName(ctx, wg.AccountListRequest{
		Search: request.Username,
		Region: request.Region,
	})
	if err != nil && !errors.Is(err, wg.ErrLimitExceeded) {
		//nolint:wrapcheck
		return nil, err
	}

	if listEntry == nil {
		return nil, models.ErrNotFound
	}

	info, err := s.wg.AccountInfo(ctx, wg.AccountInfoRequest{
		Region: request.Region,
		IDs:    []int64{listEntry.AccountID},
	})
	if err != nil {
		//nolint:wrapcheck
		return nil, err
	}

	account := info[listEntry.AccountID]
	if account == nil {
		return nil, models.ErrNotFound
	}

	return account, nil
}
